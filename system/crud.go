package system

import (
	"encoding/json"
	"fmt"
	"jws/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func params(r *http.Request) (interface{}, interface{}, string, int) {
	obj, slice := model.ObjModel(mux.Vars(r)["type"])
	ty := mux.Vars(r)["type"]
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	return obj, slice, ty, id
}

const RECORDS_PER_PAGE = 50
const MAX_RECORDS_PER_PAGE = 250

func View(w http.ResponseWriter, r *http.Request) {
	obj, slice, _, _ := params(r)
	if obj == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	innerDb := DB

	if ids := r.URL.Query()["id"]; len(ids) > 0 {
		innerDb = innerDb.Or(ids)
	}

	innerDb = model.Search(innerDb, obj, r.URL.Query())
	innerDb = model.Order(innerDb, obj, r.URL.Query())

	records, err := strconv.Atoi(r.URL.Query().Get("records"))
	if err != nil || records < 1 || records > MAX_RECORDS_PER_PAGE {
		records = RECORDS_PER_PAGE
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

func Create(w http.ResponseWriter, r *http.Request) {
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

	res := DB.Create(obj)
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

func Retrieve(w http.ResponseWriter, r *http.Request) {
	obj, _, _, id := params(r)
	if obj == nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := DB.First(obj, id)
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

func Update(w http.ResponseWriter, r *http.Request) {
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

	res := DB.Save(obj)
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

func Delete(w http.ResponseWriter, r *http.Request) {
	obj, _, _, id := params(r)
	if obj == nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := DB.Delete(obj, id)
	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
