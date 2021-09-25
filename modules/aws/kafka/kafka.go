package kafka

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	kafka "github.com/aws/aws-sdk-go/service/kafka"
)

// Kafka ...
type Kafka struct {
	s       *session.Session
	cluster *string
}

// IKafka ...
type IKafka interface {
	GetBootstrapBrokers() (*kafka.GetBootstrapBrokersOutput, error)
}

// New ...
func New() (i *IKafka, err error) {
	var (
		awssession *session.Session
	)

	Cluster := os.Getenv("AWS_CLUSTER")
	AccessKey := os.Getenv("AWS_KAFKA_ACCESS_KEY_ID")
	AccessSecret := os.Getenv("AWS_KAFKA_SECRET_ACCESS_KEY")
	Region := os.Getenv("AWS_KAFKA_DEFAULT_REGION")

	if AccessKey != "" && AccessSecret != "" {
		awssession, err = session.NewSession(&aws.Config{
			Region:      aws.String(Region),
			Credentials: credentials.NewStaticCredentials(AccessKey, AccessSecret, ""),
		})
	} else {
		awssession, err = session.NewSession(&aws.Config{
			Region: aws.String(Region),
		})
	}

	if err != nil {
		i := IKafka(&Kafka{})
		return &i, err
	}

	s := IKafka(&Kafka{
		s:       awssession,
		cluster: aws.String(Cluster),
	})

	i = &s

	return
}

// GetBootstrapBrokers return kafka brokers
func (me *Kafka) GetBootstrapBrokers() (*kafka.GetBootstrapBrokersOutput, error) {
	request := kafka.GetBootstrapBrokersInput{}
	request.ClusterArn = me.cluster

	return kafka.New(me.s).GetBootstrapBrokers(&request)
}

// Get ...
func Get() (string, error) {
	svc, err := New()

	if err != nil {
		return "", err
	}

	r, err := (*svc).GetBootstrapBrokers()
	if err != nil {
		return "", err
	}

	return *r.BootstrapBrokerString, nil
}
