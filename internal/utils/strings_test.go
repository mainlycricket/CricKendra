package utils

import (
	"reflect"
	"testing"
)

func TestAddPrefixSuffix(t *testing.T) {
	type args struct {
		arr      []string
		prefixes []string
		suffixes []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "prefix & suffix both",
			args: args{
				arr:      []string{"word1", "word2", "word3"},
				prefixes: []string{"pre1", "pre2", "pre3"},
				suffixes: []string{"suf1", "suf2", "suf3"},
			},
			want: []string{"pre1word1suf1", "pre2word2suf2", "pre3word3suf3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddPrefixSuffix(tt.args.arr, tt.args.prefixes, tt.args.suffixes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddPrefixSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbersAsStrings(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "2 to 5",
			args: args{start: 2, end: 5},
			want: []string{"2", "3", "4", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumbersAsStrings(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumbersAsStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepeatWord(t *testing.T) {
	type args struct {
		word string
		n    int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1 repetition",
			args: args{word: "word", n: 1},
			want: []string{"word"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepeatWord(tt.args.word, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepeatWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
