package dto

type ParentResponse struct {
	MessageCode string `json:"mc"`
	Message     string `json:"m"`
}

type SuccessResponse[T any] struct {
	ParentResponse
	Data T `json:"lc"`
}

type ErrorResponse struct {
	ParentResponse
	Error []Error `json:"err"`
}

type Error struct {
	DebugId      string `json:"di"`
	ErrorMessage string `json:"em"`
}
