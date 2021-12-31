package verihubs

import "encoding/json"

const (
	V1Host = "https://api.verihubs.com/v1/"

	HeaderAccept      = "Accept"
	HeaderContentType = "Content-Type"
	HeaderJsonValue   = "application/json"
	HeaderAppID       = "App-ID"
	HeaderApiKey      = "API-Key"
)

// Request data struct for request send
type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    interface{}
	Queries map[string]string
}

// Basic response basic for every result
type Basic struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (b *Basic) Error() string {
	message, _ := json.Marshal(b)
	return string(message)
}
