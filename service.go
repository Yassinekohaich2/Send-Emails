package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Service struct {
	store Storage
}

func NewService(store Storage) *Service {
	return &Service{store: store}
}
func (s *Service) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/users/{userID}parse-kindle-file", s.handleParseKindleFile).Methods("POST")
	router.HandleFunc("/cloud/send-daily-insights", s.handleSendDailyInsights).Methods("GET")
}

func (s *Service) handleParseKindleFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	file, _, err := r.FormFile("file")
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing file: %v", err))
		return
	}

	defer file.Close()

	//parse that multipart file
	raw, err := parseKindleExtractFile(file)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing file: %v", err))
		return
	}
	userIDint, _ := strconv.Atoi(userID)
	if err := s.createDataFromRawBook(raw, userIDint); err != nil {
		WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error creating data from raw book: %v", err))
		return
	}

	WriteJSON(w, http.StatusOK, "Successfully parsed file")
}

func (s *Service) handleSendDailyInsights(w http.ResponseWriter, r *http.Request) {

}

func parseKindleExtractFile(file multipart.File) (*RawExtractBook, error) {
	decoder := json.NewDecoder(file)
	raw := new(RawExtractBook)
	if err := decoder.Decode(raw); err != nil {
		return nil, err
	}

	return raw, nil
}

func (s *Service) createDataFromRawBook(raw *RawExtractBook, userID int) error {

	_, err := s.store.GetBookByISBN(raw.ASIN)
	if err != nil {
		s.store.CreateBook(Book{
			ISBN:    raw.ASIN,
			Title:   raw.Title,
			Authors: raw.Authors,
		})
	}
	// create highlights
	hs := make([]Highlight, len(raw.Highlights))

	for i, h := range raw.Highlights {
		hs[i] = Highlight{
			UserID:   userID,
			BookID:   raw.ASIN,
			Text:     h.Text,
			Location: h.Location.URL,
			Note:     h.Note,
		}
	}

	err = s.store.CreateHightLights(hs)
	if err != nil {
		log.Println("Error creating highlights: ", err)
		return err
	}
	return nil

}

func GoToEmails() error {
	return nil
}
