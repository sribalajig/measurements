package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hpi/measurement/pkg/datastore"
	"github.com/hpi/measurement/pkg/model"
	"github.com/hpi/measurement/pkg/service"
)

// Server - contains a method to bootstrap a http server
type Server struct {
	measurementService *service.MeasurementService
}

// NewServer return a pointer to Server
func NewServer(dp datastore.Provider) *Server {
	return &Server{
		measurementService: service.NewMeasurementService(dp),
	}
}

// Start - starts the http server
func (server *Server) Start(port string, ch chan error) {
	r := mux.NewRouter()

	s := r.PathPrefix("/users/{userID}").Subrouter()

	s.HandleFunc("/bodyMeasurements", server.createHandler).Methods(http.MethodPost)
	s.HandleFunc("/bodyMeasurements", server.getHandler).Methods(http.MethodGet)

	log.Printf("HTTP server listening on port : %s\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)

	ch <- err
}

// createHandler handles PUT requests to /users/{userID}/bodyMeasurements
func (server *Server) createHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUserID(mux.Vars(r))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userMeasurements model.UserMeasurements

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&userMeasurements)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	userMeasurements.ID = userID

	err = server.measurementService.Create(&userMeasurements)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// getHandler handles GET request to /users/{userID}/bodyMeasurements
func (server *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUserID(mux.Vars(r))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	measurementIDs, err := parseMeasurementIDs(r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	measurements := server.measurementService.Get(userID, measurementIDs)

	if measurements == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(measurements)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func parseUserID(params map[string]string) (int, error) {
	if u, ok := params["userID"]; !ok {
		return 0, errors.New("No user id in params")
	} else if userID, err := strconv.Atoi(u); err != nil {
		return 0, err
	} else {
		return userID, nil
	}
}

func parseMeasurementIDs(query map[string][]string) ([]int, error) {
	var measurementIDs []int

	if ids, ok := query["id"]; ok {
		for _, id := range ids {
			parsedID, err := strconv.Atoi(id)
			if err != nil {
				return nil, err
			}

			measurementIDs = append(measurementIDs, parsedID)
		}
	}

	return measurementIDs, nil
}
