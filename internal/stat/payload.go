package stat

type GetStatResponce struct {
	Period string `json:"period"`
	Sum    int64  `json:"sum"`
}
