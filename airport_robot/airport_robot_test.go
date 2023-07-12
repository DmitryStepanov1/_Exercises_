package airportrobot

import "testing"

func TestSayHello_Italian(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		want     string
	}{
		{
			testName: "name without spaces",
			name:     "Petya",
			want:     "I can speak Italian: Ciao Petya!",
		},
		{
			testName: "full name",
			name:     "Dmitry Moriarti",
			want:     "I can speak Italian: Ciao Dmitry Moriarti!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := SayHello(tt.name, Italian{}); got != tt.want {
				t.Errorf("SayHello(%q, \"Italian{}\") = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestSayHello_Portuguese(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		want     string
	}{
		{
			testName: "name without spaces",
			name:     "Fabrício",
			want:     "I can speak Portuguese: Olá Fabrício!",
		},
		{
			testName: "full name",
			name:     "Manuela Alberto",
			want:     "I can speak Portuguese: Olá Manuela Alberto!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := SayHello(tt.name, Portuguese{}); got != tt.want {
				t.Errorf("SayHello(%q, \"Portuguese{}\") = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
