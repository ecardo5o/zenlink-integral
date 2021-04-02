package models

type UrlResponse struct {
	*BasicResponse
	Result string `json:"result"`
}

func GenerateUrlResponse (balance string) UrlResponse {
	return UrlResponse{
		BasicResponse: new(BasicResponse).Generate(),
		Result:        balance,
	}
}