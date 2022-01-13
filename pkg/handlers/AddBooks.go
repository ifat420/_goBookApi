package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ifat420/_goBookApi/pkg/mocks"
	"github.com/ifat420/_goBookApi/pkg/models"
)

func AddBooks(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	book.Id = 15
	// log.Fatalln(book)
	mocks.Books = append(mocks.Books, book)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mocks.Books)
}
