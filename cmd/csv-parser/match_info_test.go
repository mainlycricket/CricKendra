package main

import (
	"reflect"
	"testing"
)

func Test_getTourHostNations(t *testing.T) {
	type args struct {
		eventName string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "single_host_nation",
			args: args{eventName: "India tour of Sri Lanka"},
			want: []string{"Sri Lanka"},
		},
		{
			name: "multiple_host_nations",
			args: args{eventName: "India tour of England and Ireland"},
			want: []string{"England", "Ireland"},
		},
		{
			name: "invalid",
			args: args{eventName: "India v New Zealand Test series"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTourHostNations(tt.args.eventName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractParts() got1 = %v, want %v", got, tt.want)
			}
		})
	}
}
