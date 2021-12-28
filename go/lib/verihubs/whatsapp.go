package verihubs

import (
	"context"
	"net/http"
)

// SendWhatsAppMessage call send whatsApp message verihubs api
func (c *Client) SendWhatsAppMessage(ctx context.Context, req SendWhatsAppMessageRequest) (SendWhatsAppMessageResponse, int, error) {
	var res SendWhatsAppMessageResponse

	code, err := c.Send(ctx, Request{Method: http.MethodPost, Path: V1WhatsAppMessageSend, Body: req}, &res)
	return res, code, err
}
