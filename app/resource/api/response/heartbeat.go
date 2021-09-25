package response

type TimeStampResponse struct {
	DbTimestamp     string `json:"db_time_stamp,omitempty"`
	SystemTimeStamp string `json:"sys_time_stamp,omitempty"`
}
