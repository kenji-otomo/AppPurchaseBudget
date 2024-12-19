package app

type AppRequest struct {
	Name string `json:"name"`
}

type UpdateAppRequest struct {
	ID   *int64 `json:"id"`
	Name string `json:"name"`
}
