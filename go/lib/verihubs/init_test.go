package verihubs

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"reflect"
	"testing"
	"time"

	"gopkg.in/h2non/gock.v1"
)

var testClient *Client

func TestMain(m *testing.M) {
	log.Println("new mock client")
	testClient = New(Config{
		Host:   "https://testverihubs.com/",
		AppID:  "123",
		APIKey: "234",
	})
	code := m.Run()
	os.Exit(code)
}

func mockHttp(headers map[string]string) *gock.Request {
	matchHeaders := map[string]string{
		HeaderAccept:      HeaderJsonValue,
		HeaderContentType: HeaderJsonValue,
		HeaderAppID:       "123",
		HeaderApiKey:      "234",
	}

	for key, val := range headers {
		matchHeaders[key] = val
	}

	testGock := gock.New("https://testverihubs.com/").
		MatchHeaders(matchHeaders)

	return testGock
}

func TestNew(t *testing.T) {
	timeOut := 10 * time.Second
	transport := http.NewFileTransport(http.Dir("/"))
	cookieJar, _ := cookiejar.New(nil)

	type args struct {
		config Config
		opts   []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				config: Config{
					Host:    "https://test.com",
					AppID:   "xyz",
					APIKey:  "yyz",
					Timeout: timeOut,
				},
				opts: []Option{
					WithTimeout(timeOut),
					WithTransport(transport),
					WithCookieJar(cookieJar),
				},
			},
			want: &Client{
				http: &http.Client{
					Transport: transport,
					Jar:       cookieJar,
					Timeout:   10 * time.Second,
				},
				host:   "https://test.com",
				apiKey: "yyz",
				appID:  "xyz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.config, tt.args.opts...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
