package va_creation

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/model"
	"encoding/xml"
	"github.com/beevik/etree"
)

const (
	XMLNS          = "http://www.kotak.com/Schemas/virtualAccountMasters.xsd"
	SchemaLocation = "http://www.kotak.com/Schemas/virtualAccountMasters.xsd schema.xsd"
	Xsi            = "http://www.w3.org/2001/XMLSchema-instance"
)

type VirtualAccountMasters struct {
	XMLName           xml.Name          `xml:"virtualAccountMasters"`
	Text              string            `xml:",chardata"`
	Xmlns             string            `xml:"xmlns,attr"`
	SchemaLocation    string            `xml:"xsi:schemaLocation,attr"`
	Xsi               string            `xml:"xmlns:xsi,attr"`
	RequestParameters RequestParameters `xml:"requestParameters"`
}

type RequestParameters struct {
	Text             string `xml:",chardata"`
	SRCAPPCD         string `xml:"SRC_APP_CD"`
	UNIQUEREFNO      string `xml:"UNIQUE_REF_NO"`
	MODE             string `xml:"MODE"`
	SOURCE           string `xml:"SOURCE"`
	LOGINID          string `xml:"LOGIN_ID"`
	COMPANYCODE      string `xml:"COMPANY_CODE"`
	APAC             string `xml:"APAC"`
	MAPVIRTAPAC      string `xml:"MAP_VIRT_APAC"`
	DEALERCODE       string `xml:"DEALER_CODE"`
	DEALERNAME       string `xml:"DEALER_NAME"`
	DRNARRATION      string `xml:"DR_NARRATION"`
	CRNARRATION      string `xml:"CR_NARRATION"`
	EMAILID          string `xml:"EMAIL_ID"`
	MOBILE           string `xml:"MOBILE"`
	REMITAPACCHECK   string `xml:"REMIT_APAC_CHECK"`
	PARTIALCHECKFLAG string `xml:"PARTIAL_CHECK_FLAG"`
	IMPSENABLEDFLAG  string `xml:"IMPS_ENABLED_FLAG"`
	PRODUCTCD        string `xml:"PRODUCT_CD"`
	REGIONCD         string `xml:"REGION_CD"`
	LOCATIONCD       string `xml:"LOCATION_CD"`
	REMARKS          string `xml:"REMARKS"`
	OTHERINFO1       string `xml:"OTHER_INFO1"`
	OTHERINFO2       string `xml:"OTHER_INFO2"`
	OTHERINFO3       string `xml:"OTHER_INFO3"`
	OTHERINFO4       string `xml:"OTHER_INFO4"`
	OTHERINFO5       string `xml:"OTHER_INFO5"`
	OTHERINFO6       string `xml:"OTHER_INFO6"`
	OTHERINFO7       string `xml:"OTHER_INFO7"`
	OTHERINFO8       string `xml:"OTHER_INFO8"`
	OTHERINFO9       string `xml:"OTHER_INFO9"`
	OTHERINFO10      string `xml:"OTHER_INFO10"`
	OTHERINFO11      string `xml:"OTHER_INFO11"`
	OTHERINFO12      string `xml:"OTHER_INFO12"`
	OTHERINFO13      string `xml:"OTHER_INFO13"`
	OTHERINFO14      string `xml:"OTHER_INFO14"`
	OTHERINFO15      string `xml:"OTHER_INFO15"`
	OTHERINFO16      string `xml:"OTHER_INFO16"`
	OTHERINFO17      string `xml:"OTHER_INFO17"`
	OTHERINFO18      string `xml:"OTHER_INFO18"`
	OTHERINFO19      string `xml:"OTHER_INFO19"`
	OTHERINFO20      string `xml:"OTHER_INFO20"`
	ADDRLINE1        string `xml:"ADDR_LINE1"`
	ADDRLINE2        string `xml:"ADDR_LINE2"`
	ADDRLINE3        string `xml:"ADDR_LINE3"`
	ADDRLINE4        string `xml:"ADDR_LINE4"`
	CITY             string `xml:"CITY"`
	PINCODE          string `xml:"PIN_CODE"`
	DEACTIVEFLAG     string `xml:"DEACTIVE_FLAG"`
	CLIENTCODE       string `xml:"CLIENT_CODE"`
}

func FormatRequest(param map[string]string, va_config *model.KotakVAConfig) (vareq VirtualAccountMasters) {
	vareq.Xmlns = XMLNS
	vareq.SchemaLocation = SchemaLocation
	vareq.Xsi = Xsi
	vareq.RequestParameters.SRCAPPCD = va_config.SourceAppCd.String
	vareq.RequestParameters.UNIQUEREFNO = param["ref_no"]
	vareq.RequestParameters.MODE = va_config.Mode.String
	vareq.RequestParameters.SOURCE = va_config.Source.String
	vareq.RequestParameters.LOGINID = va_config.LoginId.String
	vareq.RequestParameters.COMPANYCODE = va_config.CompanyCode.String
	vareq.RequestParameters.APAC = va_config.Apac.String
	vareq.RequestParameters.MAPVIRTAPAC = param["virt_apac"]
	vareq.RequestParameters.DEALERNAME = param["name"]
	vareq.RequestParameters.REMITAPACCHECK = va_config.RemitApacFlag.String
	vareq.RequestParameters.PARTIALCHECKFLAG = va_config.PartialCheckFlag.String
	vareq.RequestParameters.IMPSENABLEDFLAG = va_config.IsIMPSEnabled.String

	return vareq
}

func AppendXmlVersion(msg string) (byt []byte) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="utf-8"`)
	doc.ReadFromString(msg)
	doc.Indent(etree.NoIndent)

	str, err := doc.WriteToString()
	if err != nil {
		panic(err)
	}

	byt = []byte(str)
	return byt

}
