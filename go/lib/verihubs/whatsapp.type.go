package verihubs

// SendWhatsAppMessageRequest request message for send whatsApp message
type SendWhatsAppMessageRequest struct {
	Msisdn      string   `json:"msisdn"`
	Content     []string `json:"content,omitempty"`
	LangCode    string   `json:"lang_code,omitempty"`
	CallbackURL string   `json:"callback_url,omitempty"`
}

// SendWhatsAppMessageResponse response data for send whatsApp message
type SendWhatsAppMessageResponse struct {
	Msisdn    string `json:"msisdn"`
	SessionID string `json:"session_id"`
	Basic
}
