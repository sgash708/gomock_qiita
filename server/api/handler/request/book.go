package request

type GetBookRequest struct {
	UUID string `param:"uuid"`
}

type CreateBookRequest struct {
	Name string `json:"name"`
}

type UpdateBookRequest struct {
	UUID string `param:"uuid"`
	Name string `json:"name"`
}

type DeleteBookRequest struct {
	UUID string `param:"uuid"`
}
