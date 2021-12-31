package verihubs

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_SendWhatsAppMessage(t *testing.T) {
	type args struct {
		req SendWhatsAppMessageRequest
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(a args)
		want    SendWhatsAppMessageResponse
		want1   int
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				req: SendWhatsAppMessageRequest{
					Msisdn:      "123",
					Content:     []string{"a"},
					LangCode:    "en",
					CallbackURL: "https://abc.com/www",
				},
			},
			mockFn: func(a args) {
				mockHttp(nil).
					Post(V1WhatsAppMessageSend).
					JSON(a.req).
					Reply(http.StatusCreated).
					BodyString(`{"msisdn":"123","session_id":"1"}`)
			},
			want: SendWhatsAppMessageResponse{
				Msisdn:    "123",
				SessionID: "1",
			},
			want1: http.StatusCreated,
		},
		{
			name: "got err from client",
			args: args{
				req: SendWhatsAppMessageRequest{
					Msisdn:      "666",
					Content:     []string{"a"},
					LangCode:    "en",
					CallbackURL: "https://abc.com/www",
				},
			},
			mockFn: func(a args) {
				mockHttp(nil).
					Post(V1WhatsAppMessageSend).
					JSON(a.req).
					Reply(http.StatusBadRequest).
					BodyString(`{"code":400,"message":"Send","status":1}`)
			},
			want:    SendWhatsAppMessageResponse{},
			want1:   http.StatusBadRequest,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClient

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			got, got1, err := c.SendWhatsAppMessage(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendWhatsAppMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendWhatsAppMessage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SendWhatsAppMessage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
