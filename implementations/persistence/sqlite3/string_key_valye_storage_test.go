package sqlite3_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	. "xeronith/url-shortener/contracts/caching"
	. "xeronith/url-shortener/contracts/persistence"
	"xeronith/url-shortener/implementations/caching"
	"xeronith/url-shortener/implementations/persistence"
)

var (
	cache   StringKeyValueCache
	storage StringKeyValueStorage
)

func TestMain(m *testing.M) {
	cache = caching.CreateStringKeyValueCache(InMemoryStringKeyValueCache)
	storage = persistence.CreateStringKeyValueStorage(Sqlite3StringKeyValueStorage, cache)

	exitCode := m.Run()
	storage.Destroy()
	os.Exit(exitCode)
}

func Test_StringKeyValueStorage_Persist(t *testing.T) {
	type args struct {
		key   string
		value string
	}

	unixNano := time.Now().UnixNano()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Case1",
			args{
				key:   fmt.Sprintf("Key%d", unixNano),
				value: fmt.Sprintf("Value%d", unixNano),
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := storage.Persist(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Persist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Benchmark_StringKeyValueStorage_Persist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := storage.Persist(fmt.Sprintf("Key%d", time.Now().UnixNano()), "Value"); err != nil {
			b.Fatal(err)
		}
	}
}

func Test_StringKeyValueStorage_Retrieve(t *testing.T) {
	type args struct {
		key string
	}

	unixNano := time.Now().UnixNano()
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Case1",
			args{
				fmt.Sprintf("Key%d", unixNano),
			},
			fmt.Sprintf("Value%d", unixNano),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := storage.Persist(tt.args.key, tt.want); (err != nil) != tt.wantErr {
				t.Errorf("Persist() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := storage.Retrieve(tt.args.key, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Retrieve() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
