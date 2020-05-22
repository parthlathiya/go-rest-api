package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// User 3 attributes email, first_name and last_name
type User struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", healthCheck).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user", getUser).Methods("GET").Queries("email", "{email}")

	log.Fatal(http.ListenAndServe(":8092", router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK")
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "test_user"
	dbPass := "test_password"
	dbName := "boilerplate"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Fprintf(w, "Something went wrong!!")
		log.Println("Error while decoding req body", err.Error())
	}

	db := dbConn()
	defer db.Close()

	insertQuery, err := db.Prepare("INSERT INTO myapp_user(email, first_name, last_name) VALUES(?,?,?)")
	if err != nil {
		fmt.Fprintf(w, "Something went wrong!!")
		log.Println("Error while creating user", err.Error())
	}
	insertQuery.Exec(newUser.Email, newUser.FirstName, newUser.LastName)

	fmt.Fprintf(w, "User created")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var email string
	email = r.URL.Query().Get("email")

	db := dbConn()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM myapp_user WHERE email=?", email)

	var id int
	newUser := User{}
	err := row.Scan(&id, &newUser.Email, &newUser.FirstName, &newUser.LastName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Fprintf(w, "Something went wrong!!")
		log.Println("Error while getting user", err.Error())
	}

	json.NewEncoder(w).Encode(newUser)
}
