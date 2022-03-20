package user

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"httpreq/pkg/init_DB"
	"io/ioutil"
	"net/http"
)

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var d *Delete
		if err := json.Unmarshal(content, &d); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		deleteID, err := primitive.ObjectIDFromHex(d.Target)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(err.Error()))
			return
		}
		result, err := init_DB.Collection.DeleteOne(init_DB.Ctx, bson.M{"_id": deleteID})
		if err != nil {
			rw.WriteHeader(http.StatusNotImplemented)
			rw.Write([]byte(err.Error()))
			return
		}
		if result.DeletedCount == 0 {
			rw.Write([]byte("Пользователя с таким ID не существует!"))
			return
		}
		filter := bson.D{{"friends", d.Target}}
		cur, err := init_DB.Collection.Find(init_DB.Ctx, filter)
		if err != nil {
			rw.WriteHeader(http.StatusNotImplemented)
			rw.Write([]byte(err.Error()))
			return
		}
		for cur.Next(init_DB.Ctx) {
			var u *User
			var i *userID
			if err := cur.Decode(&u); err != nil {
				rw.WriteHeader(http.StatusNotFound)
				rw.Write([]byte(err.Error()))
				return
			}
			if err := cur.Decode(&i); err != nil {
				rw.WriteHeader(http.StatusNotFound)
				rw.Write([]byte(err.Error()))
				return
			}
			pos := positionID(u.Friends, d.Target)
			RemoveId(u, u.Friends, pos)
			if _, err := init_DB.Collection.UpdateOne(
				init_DB.Ctx,
				bson.M{"_id": i.ID},
				bson.D{
					{"$set", bson.D{{"friends", u.Friends}}},
				},
			); err != nil {
				rw.WriteHeader(http.StatusNotImplemented)
				rw.Write([]byte(err.Error()))
				return
			}
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Пользователь с ID " + d.Target + " успешно удалён.\n"))
		return
	}
	rw.WriteHeader(http.StatusBadRequest)
}
