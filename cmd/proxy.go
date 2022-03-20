package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var counter int

func main() {

	http.HandleFunc("/create", createUserProxy)
	http.HandleFunc("/make_friends", makeFriendsProxy)
	http.HandleFunc("/user", deleteUserProxy)
	http.HandleFunc("/friends/", getFriendsProxy)
	http.HandleFunc("/", updateAgeProxy)
	log.Fatalln(http.ListenAndServe("localhost:9000", nil))

}

func createUserProxy(rw http.ResponseWriter, r *http.Request) {
	textByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	text := string(textByte)
	fmt.Println(text)

	if counter == 0 {
		resp, err := http.Post("http://localhost:8080/create", "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		textByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		defer resp.Body.Close()
		fmt.Println("1 instance: " + string(textByte))
		rw.Write([]byte("1 instance: " + string(textByte)))
		counter++
		return
	}
	resp, err := http.Post("http://localhost:8081/create", "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	textByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	fmt.Println("2 instance: " + string(textByte))
	rw.Write([]byte("2 instance: " + string(textByte)))
	counter--

}

func makeFriendsProxy(rw http.ResponseWriter, r *http.Request) {
	textByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	text := string(textByte)
	fmt.Println(text)

	if counter == 0 {
		resp, err := http.Post("http://localhost:8080/make_friends", "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		textByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		defer resp.Body.Close()
		fmt.Println("1 instance: " + string(textByte))
		rw.Write([]byte("1 instance: " + string(textByte)))
		counter++
		return
	}
	resp, err := http.Post("http://localhost:8081/make_friends", "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	textByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	fmt.Println("2 instance: " + string(textByte))
	rw.Write([]byte("2 instance: " + string(textByte)))
	counter--

}

func deleteUserProxy(rw http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	textByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	text := string(textByte)
	fmt.Println(text)

	if counter == 0 {
		request, err := http.NewRequest("DELETE", "http://localhost:8080/user", bytes.NewBuffer([]byte(text)))
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		resp, err := client.Do(request)
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		textByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		defer resp.Body.Close()
		fmt.Println("1 instance: " + string(textByte))
		rw.Write([]byte("1 instance: " + string(textByte)))
		counter++
		return
	}
	request, err := http.NewRequest("DELETE", "http://localhost:8081/user", bytes.NewBuffer([]byte(text)))
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	resp, err := client.Do(request)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	textByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	fmt.Println("2 instance: " + string(textByte))
	rw.Write([]byte("2 instance: " + string(textByte)))
	counter--

}

func getFriendsProxy(rw http.ResponseWriter, r *http.Request) {
	textByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	path := r.URL.Path
	userId := strings.TrimPrefix(path, "/friends/")
	text := string(textByte)
	fmt.Println(text)

	if counter == 0 {
		resp, err := http.Get("http://localhost:8080/friends/" + userId)
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		textByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		defer resp.Body.Close()
		fmt.Println("1 instance: " + string(textByte))
		rw.Write([]byte("1 instance: " + string(textByte)))
		counter++
		return
	}
	resp, err := http.Get("http://localhost:8081/friends/" + userId)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	textByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	fmt.Println("2 instance: " + string(textByte))
	rw.Write([]byte("2 instance: " + string(textByte)))
	counter--

}

func updateAgeProxy(rw http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	textByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	path := r.URL.Path
	userId := strings.TrimPrefix(path, "/")
	text := string(textByte)
	fmt.Println(text)

	if counter == 0 {
		request, err := http.NewRequest("PUT", "http://localhost:8080/"+userId, bytes.NewBuffer([]byte(text)))
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		resp, err := client.Do(request)
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		textByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			rw.WriteHeader(http.StatusServiceUnavailable)
			rw.Write([]byte(err.Error()))
			return
		}
		defer resp.Body.Close()
		fmt.Println("1 instance: " + string(textByte))
		rw.Write([]byte("1 instance: " + string(textByte)))
		counter++
		return
	}
	request, err := http.NewRequest("PUT", "http://localhost:8081/"+userId, bytes.NewBuffer([]byte(text)))
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	resp, err := client.Do(request)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	textByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	fmt.Println("2 instance: " + string(textByte))
	rw.Write([]byte("2 instance: " + string(textByte)))
	counter--

}


