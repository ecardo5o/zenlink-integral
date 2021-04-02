package models

type UserInfoResponse struct {
	*BasicResponse
	Result bool `json:"registerResult"`
}

func (pr *UserInfoResponse) Generate(registerResult bool) *UserInfoResponse {
	return &UserInfoResponse{
		BasicResponse: new(BasicResponse).Generate(),
		Result:        registerResult,
	}
}
