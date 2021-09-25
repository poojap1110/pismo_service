package s3

import (
	"bytes"
	"net/http"
	"os"
	"strings"
	"time"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/config"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	AwsS3 "github.com/aws/aws-sdk-go/service/s3"
)

// S3 ...
type S3 struct {
	s      *session.Session
	bucket *string
	acl    *string
}

// IUpload ...
type IUpload interface {
	UploadFile(string, string, string) error
	SetBucket(string) *IUpload
	SetACL(string) *IUpload
	GetObjectURL(string, ...string) (string, error)
}

const (
	InstanceKey = "S3Ins"
	IST         = "Asia/Kolkata"
	SGT         = "Asia/Singapore"

	fileTypeApk        = "apk"
	fileContentTypeApk = "application/vnd.android.package-archive"
)

// GetRegistry method ...
func GetRegistry() container.Registries {
	return container.Registries{
		container.Registry{
			Key: InstanceKey,
			Value: map[string]*IUpload{
				IST: nil,
				SGT: nil,
			},
		},
	}
}

// GetInstance function - initialize new connection to aws with singleton global scope.
func GetInstance(c *container.Container, tenantTimezone string) (s *IUpload, err error) {
	if s = c.Get(InstanceKey).(map[string]*IUpload)[tenantTimezone]; s != nil {
		return s, nil
	}

	awsConfs := config.GetInstance(c).GetTechnicalConfigs(constant.TechnicalAWS)
	key := awsConfs.(map[string]string)["key"]
	secret := awsConfs.(map[string]string)["secret"]
	region := awsConfs.(map[string]string)["region"]
	bucket := awsConfs.(map[string]string)["bucket"]

	if key == "" || secret == "" || region == "" || bucket == "" {
		panic("AWS config values required but not declared")
	}

	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	})

	if err != nil {
		i := IUpload(&S3{})
		return &i, err
	}

	i := IUpload(&S3{
		s:      session,
		bucket: aws.String(bucket),
	})

	s = &i
	contValues := c.Get(InstanceKey)

	if contValues != nil {
		contValues.(map[string]*IUpload)[tenantTimezone] = s
	} else {
		contValues = map[string]*IUpload{tenantTimezone: s}
	}

	c.StoreToGlobal(InstanceKey, contValues)

	return
}

// SetBucket method ...
func (me *S3) SetBucket(bucket string) *IUpload {
	me.bucket = aws.String(bucket)
	i := IUpload(me)
	return &i
}

// SetACL method ...
func (me *S3) SetACL(acl string) *IUpload {
	me.acl = aws.String(acl)
	i := IUpload(me)
	return &i
}

// UploadFile method ...
func (me *S3) UploadFile(fileDir string, cloudFileDir string, aclType string) error {
	file, err := os.Open(fileDir)

	if err != nil {
		return err
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)
	contentType := ""

	if rawFileExtension := strings.ToLower(fileInfo.Name()[len(fileInfo.Name())-3:]); rawFileExtension == fileTypeApk {
		contentType = fileContentTypeApk
	} else {
		contentType = http.DetectContentType(buffer)
	}

	acl := constant.S3ACLTypePublic

	if aclType != constant.S3ACLTypePublic {
		acl = constant.S3ACLTypePrivate
	}

	_, err = AwsS3.New(me.s).PutObject(&AwsS3.PutObjectInput{
		Bucket:               me.bucket,
		Key:                  aws.String(cloudFileDir),
		ACL:                  aws.String(acl),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(contentType),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}

// GetObjectURL method ...
// args[0] - filename
func (me *S3) GetObjectURL(key string, args ...string) (url string, err error) {
	params := &AwsS3.GetObjectInput{
		Bucket: me.bucket,
		Key:    aws.String(key),
		ResponseContentDisposition: aws.String(`attachment; filename="funny-cat.jpg"`),
	}

	// set filename
	if len(args) > 0 && args[0] != "" {
		params.SetResponseContentDisposition(`attachment; filename="`+args[0]+`"`)
	}

	req, _ := AwsS3.New(me.s).GetObjectRequest(params)

	url, err = req.Presign(604800 * time.Second)
	if err != nil {
		return
	}

	return
}

// Upload function ...
func Upload(c *container.Container, fileDir string, cloudFileDir string, timezone string, aclType string) error {
	u, err := GetInstance(c, timezone)

	if err != nil {
		return err
	}

	return (*u).UploadFile(fileDir, cloudFileDir, aclType)
}

// GetSignedURL function ...
func GetSignedURL(c *container.Container, key string, timezone string) (url string, err error) {
	u, err := GetInstance(c, timezone)

	if err != nil {
		return
	}

	return (*u).GetObjectURL(key)
}

// GetSignedURLWithModifiedFilename function ...
func GetSignedURLWithModifiedFilename(c *container.Container, key string, timezone string, filename string) (url string, err error) {
	u, err := GetInstance(c, timezone)

	if err != nil {
		return
	}

	return (*u).GetObjectURL(key, filename)
}
