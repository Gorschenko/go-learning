package links

type LinksCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinksUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

type LinksGetAllResponse struct {
	Links []*Link `json:"links"`
	Count int64   `json:"count"`
}
