package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", HTMLpagina)         // Zorgt ervoor dat het naar de HTML-pagina gaat.
	err := http.ListenAndServe(":8080", nil) // Hier wordt aangegeven dat de server moet draaien op poort 8080.
	if err != nil {
		errors("Kon niet verbinden met de poort", err, 0) // Geeft een foutmelding als de verbinding mislukt.
	}
}

func HTMLpagina(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		http.ServeFile(writer, request, "C:\\Users\\racxs\\OneDrive\\Documenten\\Infrastructure\\Semester 2\\Pop\\web.html") // Geeft het pad naar het HTML-bestand aan dat moet worden weergegeven.
	case "POST":
		if err := request.ParseForm(); err != nil {
			fmt.Fprintf(writer, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(writer, "Post from website! r.PostFrom = %v\n", request.PostForm)

		// Haal waarden op uit het formulier
		apparaatnaam := request.FormValue("apparaatnaam")
		voornaam := request.FormValue("name")
		achternaam := request.FormValue("surname")
		geldig_tot, _ := time.Parse("2006-01-02", request.FormValue("datum")) // Parse de waarde van 'datum' naar een tijd.Time object.

		// Maak verbinding met de database
		db, err := sql.Open("mysql", "root:v01LnMA1XG*T@tcp(localhost:3306)/apparaattoekenningen")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		// Voeg gegevens toe aan de database
		_, err = db.Exec("INSERT INTO toekenningen (voornaam, achternaam, apparaat, geldig_tot, toekenning_datum) VALUES (?, ?, ?, ?, NOW())", voornaam, achternaam, apparaatnaam, geldig_tot)
		if err != nil {
			fmt.Fprintf(writer, "Niet gelukt om gegevens in database te stoppen: %v", err)
			return
		}
	}
}

func errors(boodschap string, err error, exitcode int) {
	fmt.Println(boodschap, err)
	os.Exit(exitcode)
}
