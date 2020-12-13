package redis

import "testing"

func TestEncodeWordsToRedisSpec(t *testing.T) {
	type args struct {
		words        []string
		countOfWords int
	}
	tests := []struct {
		name         string
		words        []string
		countOfWords int
		want         string
	}{
		{
			"a get string",
			[]string{"GET", "NAME", "MANOJ"},
			3,
			"$3\r\nGET\r\n$4\r\nNAME\r\n$5\r\nMANOJ\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeWordsToRedisSpec(tt.words, tt.countOfWords); got != tt.want {
				t.Errorf("EncodeWordsToRedisSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}
