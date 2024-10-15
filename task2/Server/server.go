package Server

import (
	"encoding/base64"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"math/rand"
	"net/http"
	"os"
	"task2/Serialization"
	"task2/api"
	"time"
)

type ServerLogic struct{}

func (s *ServerLogic) GetVersion(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case "GET":
		file, errOpen := os.Open("api/api.openapi.yaml")
		if errOpen != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		var openapi Serialization.Openapi
		if err := yaml.NewDecoder(file).Decode(&openapi); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		var version api.Version
		openapi.Info.Version = "v" + openapi.Info.Version
		version.Version = &openapi.Info.Version
		if err := json.NewEncoder(w).Encode(&version); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (s *ServerLogic) Decode(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case "POST":
		var input api.Input
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		var output api.Output
		decodedBytes, err := base64.StdEncoding.DecodeString(*input.Input)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		output.Output = new(string)
		*output.Output = string(decodedBytes)
		if err := json.NewEncoder(w).Encode(&output); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (s *ServerLogic) RandomShit(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case "GET":
		var timeoutSeconds int = 10 + rand.Intn(10)
		select {
		case <-r.Context().Done():
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		case <-time.After(time.Duration(timeoutSeconds) * time.Second):
			if rand.Intn(2) == 1 {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
