package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"io"
	"net/http"
	"os"
)

const (
	host     = "localhost"
	port     = "8080"
	user     = "admin"
	password = "password"
	dbname   = "json"
)

func addCorsHeader(res http.ResponseWriter, req *http.Request) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	headers.Add("Access-Control-Allow-Headers", "access-control-allow-origin")
	headers.Add("Access-Control-Allow-Methods", "POST")

}

func main() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/upload", handleUpload).Methods("OPTIONS", "POST")

	httpServer := http.Server{
		Addr:    ":6060",
		Handler: router,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		return
	}

}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	addCorsHeader(w, r)
	//fmt.Println(r.Header)
	if r.Method == "OPTIONS" {

		//fmt.Println("net")
		return
	}
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Parse the multipart form containing both JSON data and file
	err := r.ParseMultipartForm(50 << 20) // Limit the size to 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Accessing the uploaded file
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Accessing JSON data from the form
	jsonData := r.FormValue("json_data")
	jsonData = jsonData[:len(jsonData)-1] + "," + "\"src\":" + " " + "\"" + handler.Filename + "\"" + "}"
	fmt.Println("JSON Data:", jsonData)
	urlExample := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	conn.Exec(context.Background(), "INSERT INTO cache (user_data) values ($1)", jsonData)

	// Create a new file to store the uploaded file
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Copy the file to the newly created file
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copying the file", http.StatusInternalServerError)
		return
	}

	fmt.Println("File uploaded successfully", handler.Filename)
	//Allow CORS here By * or specific origin

}
