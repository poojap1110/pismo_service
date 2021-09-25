package get_txn_response

import "encoding/xml"

type GetTxnResponseInXml struct {
	XMLName  xml.Name `xml:"tem:GetTxnResponseInXml"`
	InputTxn InputTxn `xml:"tem:strInput"`
}

type InputTxn struct {
	XMLName xml.Name `xml:"tem:strInput"`
	Value   string   `xml:",cdata"`
}

type PaymentEnquiry struct {
	XMLName    xml.Name `xml:"PaymentEnquiry"`
	CustomerId string   `xml:"CustomerId"`
	IBLRefNo   string   `xml:"IBLRefNo"`
}
