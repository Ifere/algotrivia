package responses

//GeneralResponse returns a standard response format
type GeneralResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	DevMessage string      `json:"dev_message,omitempty"`
	Message    string      `json:"message"`
	Status     int         `json:"status"`
}
