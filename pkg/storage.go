package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

const filename string = "passwords.txt" // eventually either prompt user for file or decide where to store

func SavePassword(app, password string) error {
	// need to check for creating the passwords file
	passwords, err := LoadPasswords()
	if err != nil {
		return err
	}
	_, ok := passwords[app]
	if ok {
		return fmt.Errorf("password already exists for %s", app)
	}
	passwords[app] = password

	data, err := json.Marshal(&passwords)
	if err != nil {
		return err
	}
	os.WriteFile("filename", data, 0600)

	return nil
}

func LoadPasswords() (map[string]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var passwords map[string]string
	if err := json.Unmarshal(data, &passwords); err != nil {
		return nil, err
	}
	return passwords, nil
}
