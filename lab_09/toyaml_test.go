package main

import "testing"

type TestServer struct {
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	Debug      bool     `json:"debug"`
	AllowedIPs []string `json:"allowed_ips"`
}

func TestToYaml(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name: "server struct",
			input: TestServer{
				Host:  "localhost",
				Port:  8080,
				Debug: true,
				AllowedIPs: []string{
					"192.168.1.1",
					"10.0.0.1",
				},
			},
			expected: `host: "localhost"
port: 8080
debug: true
allowed_ips:
  - "192.168.1.1"
  - "10.0.0.1"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToYAML(tt.input)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expected, result)
			}
		})
	}
}
