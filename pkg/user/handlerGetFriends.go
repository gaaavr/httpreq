package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"httpreq/pkg/init_DB"
	"net/http"
	"strings"
)

func GetFriends(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		path := r.URL.Path
		userId := strings.TrimPrefix(path, "/friends/")
		id, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(err.Error() + path))
			return
		}
		var u *User
		filter := bson.M{"_id": id}
		if err := init_DB.Collection.FindOne(init_DB.Ctx, filter).Decode(&u); err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(err.Error()))
			return
		}
		var userFriends string
		for i, v := range u.Friends {
			if i == len(u.Friends)-1 {
				userFriends += v + "."
				break
			}
			userFriends += v + " "

		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Список друзей " + u.Name + ": " + userFriends))
		return
	}
	rw.WriteHeader(http.StatusBadRequest)
}
