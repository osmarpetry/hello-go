package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	lightbulbs = make(map[string]bool)
)

func main() {
	lightbulbs["livingroom"] = false
	lightbulbs["kitchen"] = false

	http.HandleFunc("/healthcheck", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application-json")
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write([]byte(`{"message":"service is up and running"}`))
	})

	http.HandleFunc("/lightbulbs", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application-json")
		responseWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(responseWriter).Encode(lightbulbs)
	})

	http.HandleFunc("/lightbulbs/switch", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application-json")

		name := request.URL.Query().Get("name")
		if name == "" {
			responseWriter.WriteHeader(http.StatusBadRequest)
			responseWriter.Write([]byte(`{"message":"a light bulb name should be provided as the value of a 'name' querystring parameter"}`))
			return
		}

		if _, keyExists := lightbulbs[name]; !keyExists {
			responseWriter.WriteHeader(http.StatusNotFound)
			responseWriter.Write([]byte(`{"message":"a light bulb with the provided name doesn't exist"}`))
			return
		}

		lightbulbs[name] = !lightbulbs[name]

		responseWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(responseWriter).Encode(lightbulbs)
	})

	http.HandleFunc("/lightbulbs/create", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application-json")

		name := request.URL.Query().Get("name")
		if name == "" {
			responseWriter.WriteHeader(http.StatusBadRequest)
			responseWriter.Write([]byte(`{"message":"a light bulb name should be provided as the value of a 'name' querystring parameter"}`))
			return
		}

		if _, keyExists := lightbulbs[name]; keyExists {
			responseWriter.WriteHeader(http.StatusBadRequest)
			responseWriter.Write([]byte(`{"message":"a lightbulb with the provided name already exists"}`))
			return
		}

		lightbulbs[name] = false

		responseWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(responseWriter).Encode(lightbulbs)
	})

	http.HandleFunc("/lightbulbs/delete", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application-json")

		name := request.URL.Query().Get("name")
		if name == "" {
			responseWriter.WriteHeader(http.StatusBadRequest)
			responseWriter.Write([]byte(`{"message":"a light bulb name should be provided as the value of a 'name' querystring parameter"}`))
			return
		}

		if _, keyExists := lightbulbs[name]; !keyExists {
			responseWriter.WriteHeader(http.StatusNotFound)
			responseWriter.Write([]byte(`{"message":"a lightbulb with the provided name doesn't exist"}`))
			return
		}

		delete(lightbulbs, name)

		responseWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(responseWriter).Encode(lightbulbs)
	})

	fmt.Println("http server listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}