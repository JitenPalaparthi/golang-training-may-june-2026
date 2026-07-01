package models

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	LastModified uint64 `json:"last_modified"`
}

func NewUser(id int, name, email, status string, lastModified uint64) *User {
	return &User{id, name, email, status, lastModified}
}
