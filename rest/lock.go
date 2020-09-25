package rest

import (
	"net/http"
	"sync"
	"time"
)

const interval = time.Duration(5) * time.Minute

type lockKey struct {
	Type string
	ID   int
}

type lockValue struct {
	Token string
	Date  time.Time
}

var locks = sync.Map{}

func lockable(key lockKey, token string) bool {
	x, found := locks.Load(key)
	if !found || x.(lockValue).Token == token || x.(lockValue).Date.Add(interval).Before(time.Now()) {
		return true
	}
	return false
}

func lock(w http.ResponseWriter, r *http.Request) {
	obj, _, ty, id := params(r)
	if obj == nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := r.Header.Get("x-access-token")

	key := lockKey{Type: ty, ID: id}

	if !lockable(key, token) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	value := lockValue{Date: time.Now(), Token: token}
	locks.Store(key, value)
}

func unlock(w http.ResponseWriter, r *http.Request) {
	obj, _, ty, id := params(r)
	if obj == nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := r.Header.Get("x-access-token")

	key := lockKey{Type: ty, ID: id}

	if !lockable(key, token) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	locks.Delete(key)
}
