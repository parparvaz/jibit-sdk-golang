package jibit

func (e APIError) Error() string {
	return e.Message
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
