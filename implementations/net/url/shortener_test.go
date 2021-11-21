package url_test

import (
	"os"
	"strings"
	"testing"

	. "xeronith/url-shortener/contracts/caching"
	. "xeronith/url-shortener/contracts/net/url"
	. "xeronith/url-shortener/contracts/persistence"
	"xeronith/url-shortener/implementations/caching"
	"xeronith/url-shortener/implementations/net/url"
	"xeronith/url-shortener/implementations/persistence"
	stringsUtil "xeronith/url-shortener/utility/strings"
)

var (
	baseUrl   string
	storage   StringKeyValueStorage
	shortener Shortener
)

func TestMain(m *testing.M) {
	baseUrl = "http://localhost:8080"
	cache := caching.CreateStringKeyValueCache(InMemoryStringKeyValueCache)
	storage = persistence.CreateStringKeyValueStorage(VolatileStringKeyValueStorage, cache)
	shortener = url.NewShortener(baseUrl, storage)

	exitCode := m.Run()
	storage.Destroy()
	os.Exit(exitCode)
}

func Test_Shortener_Shorten(t *testing.T) {
	type args struct {
		url string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Case1",
			args{url: "https://www.google.com/search?client=safari&rls=en&q=sample&ie=UTF-8&oe=UTF-8"},
			false,
		},
		{
			"Case2",
			args{url: "https://google.com"},
			false,
		},
		{
			"Case3",
			args{url: "google.com"},
			true,
		},
		{
			"Case4",
			args{url: ""},
			true,
		},
		{
			"Case5",
			args{url: "   "},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shortener.Shorten(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			key := got[strings.LastIndex(got, "/")+1:]
			for i := 0; i < len(key); i++ {
				valid := false
				for j := 0; j < len(stringsUtil.AlphanumericCharacters); j++ {
					if key[i] == stringsUtil.AlphanumericCharacters[j] {
						valid = true
						break
					}
				}

				if !valid {
					t.Fatal("invalid_character")
				}
			}
		})
	}
}

func Test_Shortener_Expand(t *testing.T) {
	shortenedUrl, err := shortener.Shorten("https://google.com")
	key := shortenedUrl[strings.LastIndex(shortenedUrl, "/")+1:]
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		key string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Case1",
			args{key: "invalid_key"},
			true,
		},
		{
			"Case2",
			args{key: key},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shortener.Expand(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Expand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !shortener.UrlIsValid(got) {
				t.Errorf("Expand() got = %v, invalid_url", got)
				return
			}
		})
	}
}
