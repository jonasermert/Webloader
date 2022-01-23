package main

import "testing"

func TestValidateURL(t *testing.T) {
	test := []struct {
		s string
		want bool
	}{
		{"https://golang.org",true}
		{"golang.org",false}
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := validateURL(tt.s); got != tt.want {
				t.Error("validateURL() = %v, want %v", got, tt.want)
			}
		}
	}
}