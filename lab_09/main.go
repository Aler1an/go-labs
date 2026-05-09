package main

import "fmt"

type Server struct {
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	Debug      bool     `json:"debug"`
	AllowedIPs []string `json:"allowed_ips"`
}

func main() {
	s := Server{
		Host:  "localhost",
		Port:  8080,
		Debug: true,
		AllowedIPs: []string{
			"192.168.1.1",
			"10.0.0.1",
		},
	}

	yamlStr, err := ToYAML(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(yamlStr)
}
