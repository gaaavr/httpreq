package user

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"httpreq/pkg/init_DB"
	"io/ioutil"
	"net/http"
)

func MakeFriends(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var f *Friends
		var userOne *User
		var userTwo *User
		if err := json.Unmarshal(content, &f); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		if f.Source == f.Target {
			rw.WriteHeader(http.StatusAccepted)
			rw.Write([]byte("Нельзя добавить в друзья самого себя!"))
			return
		}
		sourceID, err := primitive.ObjectIDFromHex(f.Source)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		targetID, err := primitive.ObjectIDFromHex(f.Target)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		filterOne := bson.M{"_id": sourceID}
		if err := init_DB.Collection.FindOne(init_DB.Ctx, filterOne).Decode(&userOne); err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(err.Error()))
			return
		}
		filterTwo := bson.M{"_id": targetID}
		if err := init_DB.Collection.FindOne(init_DB.Ctx, filterTwo).Decode(&userTwo); err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(err.Error()))
			return
		}
		if check := ContainedID(userOne.Friends, f.Target); check {
			if check := ContainedID(userTwo.Friends, f.Source); check {
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte("Пользователи уже друзья!"))
				return
			}
			rw.WriteHeader(http.StatusBadGateway)
			rw.Write([]byte("Unknown error"))
			return
		}

		userOne.Friends = append(userOne.Friends, f.Target)
		userTwo.Friends = append(userTwo.Friends, f.Source)
		if _, err := init_DB.Collection.UpdateOne(
			init_DB.Ctx,
			bson.M{"_id": sourceID},
			bson.D{
				{"$set", bson.D{{"friends", userOne.Friends}}},
			},
		); err != nil {
			rw.WriteHeader(http.StatusNotImplemented)
			rw.Write([]byte(err.Error()))
			return
		}
		if _, err := init_DB.Collection.UpdateOne(
			init_DB.Ctx,
			bson.M{"_id": targetID},
			bson.D{
				{"$set", bson.D{{"friends", userTwo.Friends}}},
			},
		); err != nil {
			rw.WriteHeader(http.StatusNotImplemented)
			rw.Write([]byte(err.Error()))
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(userOne.Name + " и " + userTwo.Name + " теперь друзья."))
		return
	}

	rw.WriteHeader(http.StatusBadRequest)
}
