package inbound_reversal

type InboundRevesalResponse struct {
	CMSGenericInboundResponse CMSGenericInboundResponse `json:"CMSGenericInboundResponse"`
}

type CMSGenericInboundResponse struct {
	Header               Header               `json:"Header"`
	CMSGenericInboundRes CMSGenericInboundRes `json:"CMSGenericInboundRes"`
}

type CMSGenericInboundRes struct {
	MessageId string `json:"MessageId"`
	StatusCd  string `json:"StatusCd"`
	StatusRem string `json:"StatusRem"`
}
