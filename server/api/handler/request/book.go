package request

type GetBookRequest struct {
	UUID string `param:"uuid"`
}

type CreateBookRequst struct {
	Name string `json:"name"`
}
