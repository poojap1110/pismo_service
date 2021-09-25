package post_encrypt

type PostEncryptRequest struct {
	Payload string `json:"payload"`
	Key     string `json:"key"`
}
