package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Contact struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ContactService struct {
	Contacts map[int]Contact
}

func (c *ContactService) Create(writer http.ResponseWriter, request *http.Request) {
	var contact Contact
	err := json.NewDecoder(request.Body).Decode(&contact)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	id := len(c.Contacts) + 1
	contact.Id = id 

	c.Contacts[id] = contact 

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(contact)
	writer.WriteHeader(http.StatusCreated)
}

func (c *ContactService) List(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var contacts []Contact 

	for _, ct := range c.Contacts {
		contacts = append(contacts, ct)
	}

	json.NewEncoder(writer).Encode(contacts)
}

func (c *ContactService) Get(writer http.ResponseWriter, request *http.Request, id int) {
	writer.Header().Set("Content-Type", "application/json")
	if val, ok := c.Contacts[id]; ok {
		json.NewEncoder(writer).Encode(val)
	} else {
		http.Error(writer, "Contact not found", http.StatusNotFound)
	}

}

func (c *ContactService) Delete(writer http.ResponseWriter, request *http.Request, id int) {
	writer.Header().Set("Content-Type", "application/json")
	if _, ok := c.Contacts[id]; ok {
		delete(c.Contacts, id)
		writer.WriteHeader(http.StatusOK)
	} else {
		http.Error(writer, "Contact not found", http.StatusNotFound)
	}
}

func (c *ContactService) Update(writer http.ResponseWriter, request *http.Request, id int) {
	writer.Header().Set("Content-Type", "application/json")

	var contact Contact 
	err := json.NewDecoder(request.Body).Decode(&contact)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := c.Contacts[id]; ok {
		contact.Id = id 
		c.Contacts[id] = contact
	} else {
		http.Error(writer, "Contact not found", http.StatusNotFound)
	}
}

func handleUpdateContact(writer http.ResponseWriter, request *http.Request, service *ContactService) {
	q := request.URL.Query()
	if q.Get("id") != "" {
		id, _ := strconv.Atoi(q.Get("id"))
		service.Update(writer, request, id)
	} else {
		http.Error(writer, "Contact not found", http.StatusNotFound)
	}
}

func handleDeleteContact(writer http.ResponseWriter, request *http.Request, service *ContactService) {
	q := request.URL.Query()
	if q.Get("id") != "" {
		id, _ := strconv.Atoi(q.Get("id"))
		service.Delete(writer, request, id)
	} else {
		http.Error(writer, "Contact not found", http.StatusNotFound)
	}
}

func handleGetContacts(writer http.ResponseWriter, request *http.Request, service *ContactService) {
	q := request.URL.Query()
	if q.Get("id") != "" {
		id, _ := strconv.Atoi(q.Get("id"))
		service.Get(writer, request, id)
	} else {
		service.List(writer, request)
	}
}

func handleCreateContact(writer http.ResponseWriter, request *http.Request, service *ContactService) {
	service.Create(writer, request)
}

func main() {
	service := &ContactService{Contacts: make(map[int]Contact)}
	mux := http.NewServeMux()

	mux.HandleFunc("/contacts", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method{
		case http.MethodGet:
			handleGetContacts(writer, request, service)
		case http.MethodPost:
			handleCreateContact(writer, request, service)
		case http.MethodDelete:
			handleDeleteContact(writer, request, service)
		case http.MethodPut:
			handleUpdateContact(writer, request, service)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
		
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}