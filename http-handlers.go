package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleInsertCard(w http.ResponseWriter, r *http.Request) {
	var err error
	var newCard Card
	fmt.Println(r)
	_ = json.NewDecoder(r.Body).Decode(&newCard)

	// valid card struct, return 400 + error in body
	err = newCard.Validate()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return

	}

	title := newCard.Title
	cardType := newCard.CardType
	desc1 := newCard.DescriptionBox1
	desc2 := newCard.DescriptionBox2
	desc3 := newCard.DescriptionBox3
	prio := newCard.Priority
	sev := newCard.Severity
	assTo := newCard.AssignedTo
	caseNum := newCard.CaseNumber
	col := newCard.Column

	var uid int
	uid, err = insertCard(title, cardType, desc1, desc2, desc3, prio, sev, assTo, caseNum, col)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("failed to insert card"))
		return
	} else {
		newCard.CardID = uid
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(newCard)
	}
}

func handleDeleteCard(w http.ResponseWriter, r *http.Request) {
	var err error

	cardID := mux.Vars(r)

	fmt.Println(cardID)

	var rowsDeleted int
	rowsDeleted, err = removeCard(cardID["id"])

	w.Header().Set("Content-type", "text/plain")

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("failed to delete card"))
		return
	}

	if rowsDeleted == 1 {
		w.Write([]byte("successfully deleted"))

	} else {
		w.WriteHeader(404)
	}
}

func handleUpdateCard(w http.ResponseWriter, r *http.Request) {
	var err error
	var newCard Card

	_ = json.NewDecoder(r.Body).Decode(&newCard)
	cardid := newCard.CardID
	title := newCard.Title
	cardType := newCard.CardType
	desc1 := newCard.DescriptionBox1
	desc2 := newCard.DescriptionBox2
	desc3 := newCard.DescriptionBox3
	prio := newCard.Priority
	sev := newCard.Severity
	assTo := newCard.AssignedTo
	caseNum := newCard.CaseNumber
	col := newCard.Column

	var updated int
	updated, err = updateCard(cardid, title, cardType, desc1, desc2, desc3, prio, sev, assTo, caseNum, col)

	if err != nil {
		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(500)
		w.Write([]byte("failed to update card"))
	} else if updated == 1 {
		w.Header().Set("Content-type", "application/json")
		_ = json.NewEncoder(w).Encode(newCard)
	} else {
		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(404)

	}

}

func handleGetCards(w http.ResponseWriter, r *http.Request) {
	var err error

	yourCards := []Card{}

	yourCards, err = getCards()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("failed to get cards"))
	} else {
		w.Header().Set("Content-type", "application/json")
		_ = json.NewEncoder(w).Encode(yourCards)
	}
}

func handleGetCard(w http.ResponseWriter, r *http.Request) {
	var err error
	var yourCard Card

	urlID := mux.Vars(r)

	yourCard, err = getCard(urlID["id"])
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("failed to get card"))
		fmt.Println(err)
	} else {
		w.Header().Set("Content-type", "application/json")
		_ = json.NewEncoder(w).Encode(yourCard)
	}
}

func handlePostComment(w http.ResponseWriter, r *http.Request) {

}

func handlePutComment(w http.ResponseWriter, r *http.Request) {

}
