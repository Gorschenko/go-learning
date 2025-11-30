package links

type LinksCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinksCreateResponse struct {
	Link
}
