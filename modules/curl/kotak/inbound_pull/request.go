package inbound_pull

type InboundPullRequest struct {
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
}
