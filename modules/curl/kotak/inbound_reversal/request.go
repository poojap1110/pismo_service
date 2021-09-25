package inbound_reversal

type InboundRevesalRequest struct {
	Header               Header               `json:"Header"`
	CMSGenericInboundReq CMSGenericInboundReq `json:"CMSGenericInboundReq"`
}

type Header struct {
	SrcAppCd  string `json:"SrcAppCd"`
	RequestID string `json:"RequestID"`
}

type CMSGenericInboundReq struct {
	MerchantName   string `json:"MerchantName"`
	MerchantSecret string `json:"MerchantSecret"`
	RefNo          string `json:"RefNo"`
}
