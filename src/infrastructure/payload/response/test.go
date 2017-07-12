package response

type TestResponse struct {
	Hello string
}

func NewTestResponse(who string) *TestResponse {
	return &TestResponse{Hello: who}
}
