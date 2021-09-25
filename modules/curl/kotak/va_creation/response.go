package va_creation

import "encoding/xml"

type VirtualAccountMastersResponse struct {
	XMLName            xml.Name           `xml:"virtualAccountMasters"`
	Text               string             `xml:",chardata"`
	Xmlns              string             `xml:"xmlns,attr"`
	ResponseParameters ResponseParameters `xml:"responseParameters"`
}

type ResponseParameters struct {
	Text         string `xml:",chardata"`
	UNIQUEREFNO  string `xml:"UNIQUE_REF_NO"`
	ERRORFLAG    string `xml:"ERROR_FLAG"`
	RESDESC      string `xml:"RES_DESC"`
	PMAPVIRTAPAC string `xml:"P_MAP_VIRT_APAC"`
}
