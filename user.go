package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func GetPasswordHash(username string) (hash string, ok bool) {
	jsonFile, err := os.Open("database/users.json")

	if err != nil {
		fmt.Println("Cannot open users.json")
		return "", false
	}

	defer jsonFile.Close()

	jsonBytes, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("Cannot read users.json")
		return "", false
	}

	user := map[string]string{}

	json.Unmarshal(jsonBytes, &user)

	hash, ok = user[username]
	return hash, ok
}
