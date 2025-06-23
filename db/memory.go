package db

import (
	"go-invest-yourself/models"
	"sync"
)

var (
	Users = make(map[string]*models.User)
	mu    sync.Mutex
)

func AddUser(user *models.User) {
	mu.Lock()
	defer mu.Unlock()
	Users[user.ID] = user
}

func GetUser(id string) *models.User {
	mu.Lock()
	defer mu.Unlock()
	return Users[id]
}
