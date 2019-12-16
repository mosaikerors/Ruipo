package main

import (
	"math"
	"testing"
)

func Test_shortLinkToLongInt(t *testing.T) {
	type args struct {
		shortLink string
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test1", args{"0"}, 0},
		{"test2", args{"10"}, 64*1},
		{"test3", args{"110"}, 64*64 + 64*1},
		{"test4", args{"abc"}, 64*64*36 + 64*37 + 38},
		{"test5", args{"ABC"},64*64*10 + 64*11 + 12 },
		{"test6", args{"ABab01"},
			int64(math.Pow(64,5)*10+math.Pow(64, 4)*11+math.Pow(64,3)*36+math.Pow(64, 2)*37+1)},
		{"test7", args{"?"}, 63},
		{"test8",args{"!"} , 62},
		{"test9", args{"!?"}, 62*64 + 63}}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortLinkToLongInt(tt.args.shortLink); got != tt.want {
				t.Errorf("shortLinkToLongInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longIntToShortLink(t *testing.T) {
	type args struct {
		index int64
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{0}, "0"},
		{"test2", args{64*1}, "10"},
		{"test3", args{64*64 + 64*1}, "110"},
		{"test4", args{64*64*36 + 64*37 + 38}, "abc"},
		{"test5", args{64*64*10 + 64*11 + 12}, "ABC"},
		{"test6", args{int64(math.Pow(64,5)*10+math.Pow(64, 4)*11+math.Pow(64,3)*36+math.Pow(64, 2)*37+1)},
			"ABab01"},
		{"test7", args{63}, "?"},
		{"test8",args{62} , "!"},
		{"test9", args{62*64 + 63}, "!?"}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longIntToShortLink(tt.args.index); got != tt.want {
				t.Errorf("longIntToShortLink() = %v, want %v", got, tt.want)
			}
		})
	}
}