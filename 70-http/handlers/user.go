package handlers

import (
	"encoding/json"
	"http-demo/models"
	"io"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		slog.Warn("called on not implemented http method" + r.Method)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("oops! something went wrong, try again or consult admin"))
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.User

	err = json.Unmarshal(bytes, &user)
	if err != nil {
		w.Write([]byte("oops! something went wrong, try again or consult admin"))
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Id = rand.IntN(10000)
	user.Status = "active"
	user.LastModified = uint64(time.Now().Unix())

	bytes, err = json.Marshal(user)
	if err != nil {
		w.Write([]byte("oops! something went wrong, try again or consult admin"))
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	SaveUser("users.txt", append(bytes, byte('\n')))

	slog.Debug("ok")
	w.Write([]byte("ok"))
}

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
