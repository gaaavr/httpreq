package user

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"httpreq/pkg/init_DB"
	"io/ioutil"
	"net/http"
	"strings"
)

func AgeUpdate(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		path := r.URL.Path
		userId := strings.TrimPrefix(path, "/")
		id, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(err.Error()))
			return
		}
		var u *User
		filter := bson.M{"_id": id}
		if err := init_DB.Collection.FindOne(init_DB.Ctx, filter).Decode(&u); err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(err.Error()))
			return
		}
		var a *Age
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		if err := json.Unmarshal(content, &a); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		if _, err := init_DB.Collection.UpdateOne(
			init_DB.Ctx,
			bson.M{"_id": id},
			bson.D{
				{"$set", bson.D{{"age", a.NewAge}}},
			},
		); err != nil {
			rw.WriteHeader(http.StatusNotImplemented)
			rw.Write([]byte(err.Error()))
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Возраст пользователя " + u.Name + " был успешно изменён."))
		return
	}
	rw.WriteHeader(http.StatusBadRequest)
}
