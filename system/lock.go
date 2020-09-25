package system

import (
	"net/http"
	"sync"
	"time"
)

const interval = time.Duration(5) * time.Minute

type lockKey struct {
	Type string
	Id int
}

type lockValue struct {
	Token string
	Date time.Time
}

var locks = sync.Map{}

func lockable(key lockKey, token string) bool {
	x, found := locks.Load(key)
	if !found || x.(lockValue).Token == token || x.(lockValue).Date.Add(interval).Before(time.Now()) {
		return true
	}
	return false
}

func lock(ty string, id int, token string) bool {
	key := lockKey{Type: ty, Id: id}

	if !lockable(key, token) {
		return false
	}

	value := lockValue{Date: time.Now(), Token: token}
	locks.Store(key, value)

	return true
}

func unlock(ty string, id int, token string) bool {
	key := lockKey{Type: ty, Id: id}

	if !lockable(key, token) {
		return false
	}

	locks.Delete(key)

	return true
}

func Lock(w http.ResponseWriter, r *http.Request) {
	obj, _, ty, id := params(r)
	if obj == nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := r.Header.Get("x-access-token")

	if lock(ty, id, token) {
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}
}

func Unlock(w http.ResponseWriter, r *http.Request) {
	obj, _, ty, id := params(r)
	if obj == nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := r.Header.Get("x-access-token")

	if unlock(ty, id, token) {
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}
}