package main

import "testing"

func Test_getCommandPath(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"ls", args{"ls"}, "/usr/bin/ls", false},
		{"lsl", args{"lsl"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCommandPath(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCommandPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCommandPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
