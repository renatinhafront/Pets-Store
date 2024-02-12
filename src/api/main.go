package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-memdb"
)

var db *memdb.MemDB

func main() {
	fmt.Println("Iniciando Banco de dados")
	if err := initializeDB(); err != nil {
		log.Fatal("Falha ao inicializar banco", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/pets", Create).Methods("POST")
	router.HandleFunc("/pets/{id}", Get).Methods("GET")
	router.HandleFunc("/pets/{id}", Put).Methods("PUT")
	router.HandleFunc("/pets/{id}", Delete).Methods("DELETE")

	router.Use(addContentType)

	fmt.Println("Iniciando API Pet Store")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Create(w http.ResponseWriter, r *http.Request) {
	//declaranco variavel pet para realizar a desearilização do body enviado
	// na requisição e transformando na struct Pet
	var pet *Pet

	//Transformando (Desearizando) o body para struct pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, "Falha ao desearilizar body", http.StatusBadRequest)
		return
	}

	//Criar uma transação de escrita
	txn := db.Txn(true)

	//Inserindo no banco de dados
	if err := txn.Insert("pet", pet); err != nil {
		http.Error(w, "Falha ao inserir banco de dados", http.StatusInternalServerError)
		return
	}

	//Confirma transação no banco
	txn.Commit()

	//Retorna para requisição o código 201 para informar que foi criado com sucesso
	w.WriteHeader(http.StatusCreated)

}

// READ - GET
func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	//Criar uma transação de leitura
	txn := db.Txn(true)

	res, err := txn.First("pet", "id", id)
	if err != nil {
		http.Error(w, "Falha ao buscar no banco", http.StatusNotFound)
		return
	}

	if res == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(res.(*Pet))
}

// Update - put
func Put(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(r.Body)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(r.Body)
}

// Structs
type Pet struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
}

func addContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func initializeDB() error {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"pet": {
				Name: "pet",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
					"name": {
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"status": {
						Name:    "status",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Status"},
					},
				},
			},
		},
	}

	// Create a new data base
	dba, err := memdb.NewMemDB(schema)
	if err != nil {
		return err
	}

	// Create a write transaction
	db = dba
	return nil
}
