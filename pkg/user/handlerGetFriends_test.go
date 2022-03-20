package user

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFriends(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(GetFriends))
	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	textBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	defer resp.Body.Close()
	if string(textBytes) != "the provided hex string is not a valid ObjectID/" {
		t.Log(string(textBytes))
		t.Fail()
	}
}
