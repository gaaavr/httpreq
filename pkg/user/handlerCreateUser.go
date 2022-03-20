package user

import (
	"encoding/json"
	"httpreq/pkg/init_DB"
	"io/ioutil"
	"net/http"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}

		defer r.Body.Close()
		var u *User
		if err := json.Unmarshal(content, &u); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		if id, err1 := init_DB.Collection.InsertOne(init_DB.Ctx, u); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err1.Error()))
			return
		} else {
			rw.WriteHeader(http.StatusCreated)
			rw.Write([]byte("Пользователь с ID " + toString(id.InsertedID) + " был создан.\n"))
			return
		}
	}
	rw.WriteHeader(http.StatusBadRequest)
}
