package models

type BasicResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
}

func (b *BasicResponse) Generate() *BasicResponse {
	b.Jsonrpc = "2.0"
	b.ID = 1
	return b
}
