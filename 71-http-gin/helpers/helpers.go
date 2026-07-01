package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"http-demo/models"
	"log/slog"
	"os"
)

func SaveUser(filename string, bytes []byte) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error(err.Error())
		return err
		//os.Exit(2)
		//runtime.Goexit()
	}

	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(filename string) ([]models.User, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(">>>>>>>>>.")
		slog.Error(err.Error())
		return nil, err
	}

	fmt.Println(">>>>>>>>>.")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	users := make([]models.User, 0)
	// 4. Read line by line
	for scanner.Scan() {
		user := models.User{}
		line := scanner.Text() // Returns the line as a string without the newline character
		slog.Info(">>>>>>" + line)
		err := json.Unmarshal([]byte(line), &user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
