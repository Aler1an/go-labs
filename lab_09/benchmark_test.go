package main

import (
	"testing"
)

type BenchServer struct {
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	Debug      bool     `json:"debug"`
	AllowedIPs []string `json:"allowed_ips"`
}

var benchData = BenchServer{
	Host:  "localhost",
	Port:  8080,
	Debug: true,
	AllowedIPs: []string{
		"192.168.1.1",
		"10.0.0.1",
	},
}

func BenchmarkToYAML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ToYAML(benchData)
	}
}

func BenchmarkToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ToJSON(benchData)
	}
}
