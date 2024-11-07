package main

import (
	"reflect"
	"testing"
)

func Test_extractParts(t *testing.T) {
	type args struct {
		eventName string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []string
		want2 bool
	}{
		{
			name:  "single_host_nation",
			args:  args{eventName: "India tour of Sri Lanka"},
			want:  "India",
			want1: []string{"Sri Lanka"},
			want2: true,
		},
		{
			name:  "multiple_host_nations",
			args:  args{eventName: "India tour of England and Ireland"},
			want:  "India",
			want1: []string{"England", "Ireland"},
			want2: true,
		},
		{
			name:  "invalid",
			args:  args{eventName: "India v New Zealand Test series"},
			want:  "",
			want1: nil,
			want2: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := detectTour(tt.args.eventName)
			if got != tt.want {
				t.Errorf("extractParts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("extractParts() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("extractParts() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
