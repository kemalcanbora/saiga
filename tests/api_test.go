package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {
	values := map[string]string{"email":"kemalcanbora@gmail.com", "password":"123456"}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://127.0.0.1:8080/api/user/signup", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res)
}

func TestLoginUser(t *testing.T) {
	values := map[string]string{"email":"kemalcanbora@gmail.com", "password":"123456"}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://127.0.0.1:8080/api/user/login", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res)
}

func TestGetTasks(t *testing.T) {
	token := "blablabla"
	resp, err := http.Get("http://127.0.0.1:8080/api/tasks")
	resp.Header.Add("Bearer", token)
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res)
}

func TestGetMessages(t *testing.T) {
	id := "blablabla"
	token:= "blablabla"
	url := fmt.Sprintf("http://127.0.0.1:8080/api/chats/%s/messages",id)
	resp, err := http.Get(url)
	resp.Header.Add("Bearer", token)
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res)
}

func TestDownloadAttachments(t *testing.T) {
	id := "blablabla"
	token:= "blablabla"
	url := fmt.Sprintf("http://127.0.0.1:8080/api/attachments/%s/download",id)
	resp, err := http.Get(url)
	resp.Header.Add("Bearer", token)
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res)
}