package verihubs

import (
	"context"
	"log"
	"net/http"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestClient_Send(t *testing.T) {
	defer gock.Off()

	type req struct {
		test string
	}

	type res struct {
		Result string `json:"result"`
	}

	result := res{}

	type args struct {
		host     string
		req      Request
		response interface{}
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(a args)
		want    int
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				req: Request{
					Method: http.MethodPost,
					Path:   "send/test",
					Headers: map[string]string{
						"key-head": "val-head",
					},
					Body: req{test: "a"},
					Queries: map[string]string{
						"param": "value",
					},
				},
				response: &result,
			},
			mockFn: func(a args) {
				mockHttp(a.req.Headers).
					Post(a.req.Path).
					MatchParams(a.req.Queries).
					MatchType("json").
					JSON(a.req.Body).
					Reply(http.StatusOK).
					BodyString(`{"result":"x"}`)
			},
			want: http.StatusOK,
		},
		{
			name: "failed unmarshal",
			args: args{
				req: Request{
					Method: http.MethodPost,
					Path:   "send/test",
					Headers: map[string]string{
						"key-head": "val-head",
					},
					Body: req{test: "a"},
					Queries: map[string]string{
						"param": "value",
					},
				},
				response: "!!",
			},
			mockFn: func(a args) {
				mockHttp(a.req.Headers).
					Post(a.req.Path).
					MatchParams(a.req.Queries).
					MatchType("json").
					JSON(a.req.Body).
					Reply(http.StatusOK).
					BodyString(`{"result":"x"}`)
			},
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
		{
			name: "got err new request",
			args: args{
				host: "#",
				req: Request{
					Method: "!!!!",
					Queries: map[string]string{
						"s": "x",
					},
				},
			},
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
		{
			name: "got err on Do",
			args: args{
				req: Request{
					Method: http.MethodGet,
					Queries: map[string]string{
						"s": "x",
					},
					Body: `zz!!!`,
				},
			},
			want:    http.StatusInternalServerError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClient

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			got, err := c.Send(context.Background(), tt.args.req, tt.args.response)
			log.Println(err)
			log.Println(tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Send() got = %v, want %v", got, tt.want)
			}
		})
	}
}
