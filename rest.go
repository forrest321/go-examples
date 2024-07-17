package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// This will be used to demonstrate a basic data type
type serverData struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// This will represent what is normally a database, but could be JSON,
// CSV, XML, etc.
var dataStore map[int]*serverData

var url = "localhost"
var port = "8888"

func rest() {
	//Build dummy data
	dataStore = buildData()
	//Create router using new mux capabilities which would normally
	//be created with a package like gin
	mux := http.NewServeMux()
	mux.Handle("GET /{id}", http.HandlerFunc(handleGet))
	mux.Handle("POST /", http.HandlerFunc(handlePost))
	mux.Handle("PUT /", http.HandlerFunc(handlePut))
	mux.Handle("DELETE /{id}", http.HandlerFunc(handleDelete))
	fmt.Println("Router created")
	//Kick off the server and continue
	go listenAndServe(mux)
	//Wait a sec to allow the server to init
	time.Sleep(1 * time.Second)
	//Demonstrate usage of endpoints
	callEndpoints()
}

func callEndpoints() {
	callGet(1)
	callGet(5)
	callPost()
	callPut()
	callDelete(4)
	callDelete(6)
}

func callGet(id int) {
	fmt.Println("Calling GET endpoint")
	//This is the simplest way to call a GET in Go
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/%v", url, port, id))
	if err != nil {
		fmt.Println(fmt.Sprintf("Error calling GET: %v", err))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error calling GET: %v", err))
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Status : %s\n", resp.Status)
	} else {
		fmt.Printf("Response received: %s\n", string(body))
	}
}

func callPost() {
	fmt.Println("Getting data for id=11. This should get a 404:")
	callGet(11)
	fmt.Println("Calling POST endpoint")
	data := serverData{Id: 11, Name: "data-11", Description: "This is data 11"}
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error marshalling json: %v", err))
	}
	//This is the simplest way to POST in Go
	resp, err := http.Post(fmt.Sprintf("http://%s:%s/", url, port), "application/json", bytes.NewBuffer(d))
	if err != nil {
		fmt.Println(fmt.Sprintf("Error calling POST: %v", err))
	}
	defer resp.Body.Close()
	fmt.Printf("Response received: %s\n", resp.Status)
	fmt.Println("Getting data for id=11. This should succeed now")
	callGet(11)
}

func callPut() {
	fmt.Println("Object to update:")
	callGet(7)
	fmt.Println("Calling PUT endpoint")
	data := serverData{Id: 7, Name: "data-777", Description: "Updated data 777"}
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error marshalling json: %v", err))
	}

	//Currently, there is no shorthand for a PUT in Go,
	//so we have to create an http client
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s/%v", url, port, 7)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(d))
	if err != nil {
		fmt.Printf("Error creating PUT request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error calling PUT: %v\n", err)
	}
	fmt.Printf("Response received: %s\n", resp.Status)
	fmt.Println("Getting data for id=7. This should be updated now")
	callGet(7)
}

func callDelete(id int) {
	callGet(id)
	fmt.Println("Calling Delete endpoint")

	//Same for DELETEs, need a client
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s/%v", url, port, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Printf("Error creating DELETE request: %v\n", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error calling DELETE: %v\n", err)
	}
	fmt.Printf("Response received: %s\n", resp.Status)
	fmt.Printf("Getting data for id=%v This should get a 404\n", id)
	callGet(id)

}

func listenAndServe(mux *http.ServeMux) {
	fmt.Println("Starting server")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	//id is defined in the route stored in mux
	id := r.PathValue("id")
	if x, err := strconv.Atoi(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if d, ok := dataStore[x]; ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(d)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var data serverData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		//poorly formed object
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, ok := dataStore[data.Id]; ok {
		//id already exists, so we can't (or shouldn't) create it
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dataStore[data.Id] = &data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	var data serverData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, ok := dataStore[data.Id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//Update the "record"
	dataStore[data.Id] = &data
	w.WriteHeader(http.StatusOK)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	x, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, ok := dataStore[x]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//Delete the record
	delete(dataStore, x)
	w.WriteHeader(http.StatusOK)
}

func buildData() map[int]*serverData {
	//Create some data to use
	data := make(map[int]*serverData)
	for i := 0; i < 11; i++ {
		data[i] = &serverData{
			Id:          i,
			Name:        fmt.Sprintf("data-%v", i),
			Description: fmt.Sprintf("This is data #%v", i),
		}
	}
	return data
}
