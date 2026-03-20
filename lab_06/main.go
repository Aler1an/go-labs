package main

import (
	"fmt"
	"github.com/Aler1an/go-labs/lab_06/models"
)

func main() {
	s := models.Server{
		Host:  "localhost",
		Port:  8080,
		Debug: true,
		AllowedIPs: []string{
			"192.168.1.1",
			"10.0.0.1",
		},
	}

	// Reflect
	json1, _ := ToJSONReflect(s)
	fmt.Println("Reflect:")
	fmt.Println(json1)

	// EasyJSON
	json2, _ := ToJSONEasy(&s)
	fmt.Println("\nEasyJSON:")
	fmt.Println(json2)
}
