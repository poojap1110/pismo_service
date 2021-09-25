package getProduct

import (
	"os"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/tcp/config"
	payloadConfig "bitbucket.org/matchmove/go-payloads/config"
	"github.com/go-errors/errors"
)

const (
	// ResourceName ...
	ResourceName = "GetProduct"

	// MockProductEnable ...
	MockProductEnable = "1"
)

// ProductProperties ...
type ProductProperties struct {
	GroupName    string
	PropertyName string
	Validation   string
}

// ProductConfigs ...
type ProductConfigs struct {
	ProductCode  string
	TenantHashID string
	Products     []ProductProperties
}

// Call function ...
func Call(c *container.Container, payload string) (string, error) {
	if os.Getenv(constant.EnvMockProduct) == MockProductEnable {
		return "", nil
	}

	i := config.GetInstance(c)
	(*i).SetResourceName(ResourceName)
	(*i).SetPayload(payload)
	(*i).FormatRequest()
	err := (*i).Call()

	if err != nil {
		return "", err
	}

	(*i).FormatResponse()

	if (*i).GetFlag() != config.SuccessCode {
		return "", FormatError((*i).GetPayload())
	}

	return (*i).GetPayload(), nil
}

// FormatRequest function ...
func FormatRequest(productID string) string {
	r := payloadConfig.Product{
		ProductID: productID,
	}

	return r.Serialize()
}

// FormatResponse function ...
func FormatResponse(r string) (fr payloadConfig.Products, err error) {
	defer func() {
		if r := recover(); err != nil {
			err = errors.New(r)
		}
	}()

	fr.Unserialize(r)

	return
}

// FormatError ...
func FormatError(r string) (err error) {
	// error struct for role is payloadConfig.ConfigResponse
	var e payloadConfig.Error
	e.Unserialize(r)

	return errors.New(e.Description)
}

// RetrieveProductConfig ...
func (me *ProductConfigs) RetrieveProductConfig(fr payloadConfig.Products, userGroup string) {
	if os.Getenv(constant.EnvMockProduct) == MockProductEnable {

		prop := ProductProperties{}
		prop.GroupName = ""
		prop.PropertyName = "mobile_country_code"
		prop.Validation = `required||pattern=[0-9]+||maxlen=6`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "mobile"
		prop.Validation = `required||pattern=[0-9]+||maxlen=20`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "first_name"
		prop.Validation = `required||minlen=1||maxlen=50||pattern=[^ ][a-zA-Z ]+`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "last_name"
		prop.Validation = `required||minlen=1||maxlen=50||pattern=[^ ][a-zA-Z ]+`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "occupation"
		prop.Validation = ""
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "email"
		prop.Validation = `required||pattern=email`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "password"
		prop.Validation = `password=number|char|upper=1||minlen=8||maxlen=32`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "birthday"
		prop.Validation = `date=2006-01-02|future=0`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "middle_name"
		prop.Validation = `min=1||max=50||pattern=[^ ][a-zA-Z ]+`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "alias"
		prop.Validation = `min=2||max=25||pattern=[a-zA-Z0-9 ]+`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "preferred_name"
		prop.Validation = `min=2||max=50||pattern=[a-zA-Z0-9 ]+`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "title"
		prop.Validation = `enum=[{"code":"Mr","description":"Mister"},{"code":"Mrs","description":"Missus"},{"code":"Miss","description":"Miss"},{"code":"Dr","description":"Doctor"},{"code":"Madam","description":"Madam"}]`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "gender"
		prop.Validation = `enum=[{"code":"male","description":"Male"},{"code":"female","description":"Female"}]`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "id_type"
		prop.Validation = `enum=[{"code":"nric","description":"NRIC"},{"code":"voters_id","description":"Voter's photo ID card"},{"code":"drivers_id","description":"Driver's license"},{"code":"passport","description":"Valid passport"},{"code":"pan","description":"PAN Card"},{"code":"ration","description":"Family ration card"},{"code":"u_bill","description":"Utility bill"},{"code":"s_license","description":"shop license"},{"code":"loi","description":"Letter of introduction"},{"code":"aadhaar","description":"Aadhaar"},{"code":"epfin","description":"EP FIN No."},{"code":"spass","description":"S PASS No."},{"code":"wp","description":"Work permit"},{"code":"cmnd","description":"Chứng minh nhân dân"},{"code":"military","description":"Military Id"},{"code":"medicare_card","description":"Medicare Card"}]`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "id_number"
		prop.Validation = ""
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "id_date_expiry"
		prop.Validation = `date=2006-01-02T15:04:05-07:00|future=1`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "id_date_issue"
		prop.Validation = `date=2006-01-02T15:04:05-07:00|future=0`
		me.Products = append(me.Products, prop)

		prop.GroupName = ""
		prop.PropertyName = "product_code"
		prop.Validation = "inmmlite"
		me.Products = append(me.Products, prop)

		me.ProductCode = "inmmlite"
		me.TenantHashID = "27b0b4985c7f8d1212d9734c4c7b1260"

		return
	}

	var validation string

	if fr.ResultCount == 0 {
		return
	}

	// To check if group exist
	var count = 0
	for _, v := range fr.Products {
		if v.GroupName == userGroup && v.GroupName != "" {
			count++
		}
	}

	if count == 0 {
		me.Products = []ProductProperties{}
		return
	}

	// If group exist
	for _, v := range fr.Products {
		prop := ProductProperties{}
		validation = ""

		if v.Required == 1 {
			validation = "required" + "||"
		}

		if v.Value != "" {
			validation = validation + v.Value
		}

		if v.Property != "product_code" {
			prop.PropertyName = v.Property
		}

		prop.Validation = validation

		if v.GroupName != "" {
			if v.GroupName == userGroup {
				prop.GroupName = v.GroupName
				if v.Property == "product_code" {
					me.ProductCode = prop.Validation
				}
			}
		} else {
			// for no GroupName, add it in the final product config as long as it does not exist yet
			if v.GroupName == "" {
				if me.CheckPropertyExist(prop.PropertyName, userGroup) {
					continue
				}

				prop.GroupName = userGroup

			}
		}

		me.TenantHashID = v.TenantHashID
		me.Products = append(me.Products, prop)
	}

}

// CheckPropertyExist ...
func (me *ProductConfigs) CheckPropertyExist(propertyName string, userGroup string) bool {
	for _, v := range me.Products {
		if v.PropertyName == propertyName {
			if v.GroupName == userGroup {
				return true
			}
		}
	}

	return false
}
