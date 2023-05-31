package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	ReadWrite = os.O_RDWR
	Create    = os.O_CREATE
	Add       = os.O_APPEND
)

var (
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
)

func main() {
	//De environment variabelen voor de connectie met de database.
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")

	// Open het logbestand voor schrijven
	logFile, err := os.OpenFile("error.log", ReadWrite|Create|Add, 0666)
	if err != nil {
		fmt.Println("Kan het logbestand niet openen:", err)
		os.Exit(1)
	}
	defer logFile.Close()

	// Gebruik het logbestand voor logging
	log.SetOutput(logFile)

	http.HandleFunc("/", HTMLpagina)        // Zorgt ervoor dat het naar de HTML-pagina gaat.
	err = http.ListenAndServe(":8080", nil) // Hier wordt aangegeven dat de server moet draaien op poort 8080.
	if err != nil {
		log.Println("Kon niet verbinden met de poort:", err)
		os.Exit(3)
	}
}

func HTMLpagina(writer http.ResponseWriter, request *http.Request) {
	htmlFilePath := "web.html" //Geeft de html file aan dat moet worden weergegeven

	switch request.Method {
	case "GET":
		http.ServeFile(writer, request, htmlFilePath)

	case "POST":
		if err := request.ParseForm(); err != nil {
			log.Println("Kon formulier niet verwerken:", err)
			return
		}
		fmt.Fprintf(writer, "Post from website! r.PostFrom = %v\n", request.PostForm)

		//Haal waarden op uit het formulier
		apparaatnaam := request.FormValue("apparaatnaam")
		voornaam := request.FormValue("name")
		achternaam := request.FormValue("surname")
		platform := request.FormValue("platform")
		serienummer := request.FormValue("serienummer")
		geldig_tot, _ := time.Parse("2006-01-02", request.FormValue("datum")) //Parse de waarde van 'datum' naar een tijd.Time object.

		//Maak verbinding met de database
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)) //%s zijn de placeholders voor de strings.
		if err != nil {
			log.Println("Kon geen verbinding maken met de database:", err)
			return
		}
		defer db.Close()

		//Voeg gegevens toe aan de database
		_, err = db.Exec("INSERT INTO toekenningen (voornaam, achternaam, apparaat, geldig_tot, platform, serienummer, toekenning_datum) VALUES (?, ?, ?, ?, ?, ?, NOW())", voornaam, achternaam, apparaatnaam, geldig_tot, platform, serienummer)
		if err != nil {
			log.Printf("Niet gelukt om gegevens in database te stoppen: %v", err)
			return
		}
	}
}
