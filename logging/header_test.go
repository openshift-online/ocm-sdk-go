package logging

import "testing"

func Test_appendHeader(t *testing.T) {
	type args struct {
		l   Level
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{Debug, "test %v"}, "Debug - test %v"},
		{"2", args{Info, "test %v"}, "Info - test %v"},
		{"3", args{Warning, "test %v"}, "Warning - test %v"},
		{"4", args{Error, "test %v"}, "Error - test %v"},
		{"5", args{Fatal, "test %v"}, "Fatal - test %v"},
		{"6", args{Info, "test %v - %s - %d"}, "Info - test %v - %s - %d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendHeader(tt.args.l, tt.args.msg); got != tt.want {
				t.Errorf("appendHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
