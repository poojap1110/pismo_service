package inbound_pull

type InboundPullResponse struct {
	CMSGenericInboundResponse CMSGenericInboundResponse `json:"CMSGenericInboundResponse"`
}

type CMSGenericInboundResponse struct {
	Header               Header               `json:"Header"`
	CMSGenericInboundRes CMSGenericInboundRes `json:"CMSGenericInboundRes"`
}

type CMSGenericInboundRes struct {
	CollectionDetails CollectionDetails `json:"CollectionDetails"`
}

type CollectionDetails struct {
	CollectionDetail []Details `json:"CollectionDetail"`
}

type Details struct {
	MasterAccNo       string `json:"Master_Acc_No"`
	RemittInfo        string `json:"Remitt_Info"`
	RemitName         string `json:"Remit_Name"`
	RemitIfsc         string `json:"Remit_Ifsc"`
	REF3              string `json:"REF3"`
	Amount            string `json:"Amount"`
	TxnRefNo          string `json:"Txn_Ref_No"`
	UtrNo             string `json:"Utr_No"`
	PayMode           string `json:"Pay_Mode"`
	ECollAccNo        string `json:"E_Coll_Acc_No"`
	RemitAcNmbr       string `json:"Remit_Ac_Nmbr"`
	Creditdateandtime string `json:"Creditdateandtime"`
	REF1              string `json:"REF1"`
	REF2              string `json:"REF2"`
	TxnDate           string `json:"Txn_Date"`
	BeneCustAcname    string `json:"Bene_Cust_Acname"`
}

type InboundPullErrResponse struct {
	ErrorCode string `json:"ErrorCode"`
	ErrorDesc string `json:"ErrorDesc"`
}
