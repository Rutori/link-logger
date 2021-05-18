package parser

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetMeta(t *testing.T) {
	ogpURL, _ := url.Parse("https://ogp.me")
	ogpPreviewURL, _ := url.Parse("https://ogp.me/logo.png")
	type args struct {
		url *url.URL
	}
	tests := []struct {
		name    string
		args    args
		want    *meta
		wantErr bool
	}{
		{
			name: "ogp.me",
			args: args{
				url: ogpURL,
			},
			want: &meta{
				Title:      "The Open Graph protocol",
				PreviewURL: ogpPreviewURL,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMeta(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMeta() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMeta() got = %v, want %v", got, tt.want)
			}
		})
	}
}
