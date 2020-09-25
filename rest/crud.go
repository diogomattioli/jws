package rest

import (
	"encoding/json"
	"fmt"
	"jws/orm"
	"net/http"
	"strconv"
)

const recordsPerPage = 50
const maxRecordsPerPage = 250

func list(w http.ResponseWriter, r *http.Request) {
	obj, slice, _, _ := params(r)
	if obj == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	innerDb := orm.DB

	if ids := r.URL.Query()["id"]; len(ids) > 0 {
		innerDb = innerDb.Or(ids)
	}

	innerDb = orm.Search(innerDb, obj, r.URL.Query())
	innerDb = orm.Order(innerDb, obj, r.URL.Query())

	records, err := strconv.Atoi(r.URL.Query().Get("records"))
	if err != nil || records < 1 || records > maxRecordsPerPage {
		records = recordsPerPage

	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	res := innerDb.Offset((page - 1) * records).Limit(records).Find(slice)
	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	total := 0
	innerDb.Model(obj).Count(&total)
	pages := total / records
	if total%records > 0 {
		pages++
	}

	out := struct {
		Data  interface{} `json:"data"`
		Total int         `json:"total"`
		Pages int         `json:"pages"`
	}{slice, total, pages}

	bytes, err := json.Marshal(out)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json := string(bytes)

	fmt.Fprintf(w, json)
}

func create(w http.ResponseWriter, r *http.Request) {
	obj, _, _, _ := params(r)
	if obj == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := orm.DB.Create(obj)
	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json := string(bytes)

	fmt.Fprintf(w, json)
}

func retrieve(w http.ResponseWriter, r *http.Request) {
	obj, _, _, id := params(r)
	if obj == nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := orm.DB.First(obj, id)
	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json := string(bytes)

	fmt.Fprintf(w, json)
}

func update(w http.ResponseWriter, r *http.Request) {
	obj, _, _, _ := params(r)
	if obj == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := orm.DB.Save(obj)
	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json := string(bytes)

	fmt.Fprintf(w, json)
}

func delete(w http.ResponseWriter, r *http.Request) {
	obj, _, _, id := params(r)
	if obj == nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := orm.DB.Delete(obj, id)
	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
