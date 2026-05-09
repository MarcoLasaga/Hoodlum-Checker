package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var (
	sessions = map[string]*Session{}
	mu       sync.Mutex
)

func generateID() string {
	return strconv.FormatInt(rand.Int63(), 36)
}

func getSession(r *http.Request) (*Session, string) {
	c, err := r.Cookie("sid")
	if err != nil {
		return nil, ""
	}

	mu.Lock()
	defer mu.Unlock()

	s, ok := sessions[c.Value]
	if !ok {
		return nil, ""
	}

	return s, c.Value
}
