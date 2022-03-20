package user

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name    string   `json:"name" bson:"name"`
	Age     string   `json:"age" bson:"age"`
	Friends []string `json:"friends" bson:"friends"`
}

type Friends struct {
	Source string `json:"source_id"`
	Target string `json:"target_id"`
}

type Delete struct {
	Target string `json:"target_id"`
}

type userID struct {
	ID primitive.ObjectID `bson:"_id"`
}

type Age struct {
	NewAge string `json:"new age"`
}

func toString(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

func ContainedID(slice []string, s string) bool {
	var check bool
	for _, v := range slice {
		if v == s {
			check = true
			return check
		}
	}
	return check
}

func positionID(slice []string, s string) int {
	for i, v := range slice {
		if v == s {
			return i
		}
	}
	return -1
}

func RemoveId(u *User, slice []string, i int) {
	slice[i] = slice[len(slice)-1]
	u.Friends = slice[:len(slice)-1]
}
