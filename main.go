package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/http2"
	"gopkg.in/yaml.v3"
)

func main() {
	// bcrypt usage example
	password := "myPassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Printf("Hashed password: %s\n", string(hashedPassword))

	// http2 usage example
	client := http2.Client{}
	fmt.Printf("HTTP2 client: %v\n", client)

	// yaml usage example
	yamlString := `
		name: John Doe
		age: 30
	`
	var data map[string]interface{}
	_ = yaml.Unmarshal([]byte(yamlString), &data)
	fmt.Printf("Unmarshalled YAML data: %v\n", data)
}
