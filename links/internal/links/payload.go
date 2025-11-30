package links

type LinksCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinksUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}
