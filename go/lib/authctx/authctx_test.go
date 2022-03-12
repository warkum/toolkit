package authctx

import "testing"

func TestAUTH_LEVEL_IsAuthorizedAs(t *testing.T) {
	type args struct {
		minLevel AUTH_LEVEL
	}
	tests := []struct {
		name string
		l    AUTH_LEVEL
		args args
		want bool
	}{
		{
			name: "SUPER authorized as ADMIN",
			l:    AUTH_SUPER,
			args: args{
				minLevel: AUTH_ADMIN,
			},
			want: true,
		},
		{
			name: "ADMIN authorized as ADMIN",
			l:    AUTH_ADMIN,
			args: args{
				minLevel: AUTH_ADMIN,
			},
			want: true,
		},
		{
			name: "BUYER unauthorized as ADMIN",
			l:    AUTH_BUYER,
			args: args{
				minLevel: AUTH_ADMIN,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsAuthorizedAs(tt.args.minLevel); got != tt.want {
				t.Errorf("AUTH_LEVEL.IsAuthorizedAs() = %v, want %v", got, tt.want)
			}
		})
	}
}
