package gqlclient

// Config define config data for gql client
type Config struct {
	Address string            `json:"address" yaml:"address"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}

// Request data that define request for gql client
type Request struct {
	Message   string
	Variables map[string]interface{}
	Headers   map[string]string
}
