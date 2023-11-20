package types

type CustomResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
}
