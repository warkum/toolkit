package httpclient

import (
	"bytes"
	"io"
	"strings"
)

func convertToReader(in interface{}) io.Reader {
	switch v := in.(type) {
	default:
		return nil
	case []byte:
		return bytes.NewReader(v)
	case string:
		return strings.NewReader(v)
	}
}
