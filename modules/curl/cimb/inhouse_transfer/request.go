package inhouse_transfer

type PostInhouseTransferRequest struct {
	ReqId    string `json:"ReqId"`
	DbtAcct  string `json:"DbtAcct"`
	CrdtAcct string `json:"CrdtAcct"`
	TrfType  string `json:"TrfType"`
	TransAmt string `json:"TransAmt"`
	Ccy      string `json:"Ccy"`
	CustRef  string `json:"CustRef"`
}
