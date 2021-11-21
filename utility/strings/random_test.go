package strings

import (
	"math/rand"
	"testing"
)

func Test_GenerateRandomString(t *testing.T) {
	type args struct {
		length int
	}

	desiredLength := 8
	tests := []struct {
		name string
		args struct {
			length int
		}
		want int
	}{
		{
			"Case1",
			args{
				length: desiredLength,
			},
			desiredLength,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomString(tt.args.length); len(got) != tt.want {
				t.Errorf("GenerateRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_GenerateRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomString(rand.Intn(4) + 4)
	}
}
