package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}
type GetAllLinksResponce struct {
	Links []Link `json:"links"`
	Count int64  `json:"count"`
}
type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}
