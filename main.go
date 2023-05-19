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
	http.HandleFunc("/", HTMLpagina)         //Zorgt ervoor dat het naar de HTMLpagina gaat.
	err := http.ListenAndServe(":8080", nil) //Hier staat dat het moet draaien op port 8080.
	if err != nil {
		errors("Kon niet verbinden met de port", err, 0) //Nog niet zeker of dit nut heeft.
	}

	// Maak een databaseverbinding
}

func HTMLpagina(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		http.ServeFile(writer, request, "C:\\Users\\racxs\\OneDrive\\Documenten\\Infrastructure\\Semester 2\\Pop\\web.html") //Dit is de html bestand.
	case "POST":
		if err := request.ParseForm(); err != nil {
			fmt.Fprintf(writer, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(writer, "Post from website! r.PostFrom = %v\n", request.PostForm)
		apparaatnaam := request.FormValue("apparaatnaam")
		voornaam := request.FormValue("name")
		achternaam := request.FormValue("surname")
		//geldig_tot := request.FormValue("toekenningdatum")
		geldig_tot, _ := time.Parse("2006-01-02", request.FormValue("datum"))

		// if geldig_tot == "" {
		// 	fmt.Fprintln(writer, "Toekenning tot is vereist")
		// 	return
		// }

		db, err := sql.Open("mysql", "root:v01LnMA1XG*T@tcp(localhost:3306)/apparaattoekenningen")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()
		_, err = db.Exec("INSERT INTO toekenningen (voornaam, achternaam, apparaat, geldig_tot, toekenning_datum) VALUES (?, ?, ?, ?, NOW())", voornaam, achternaam, apparaatnaam, geldig_tot)
		if err != nil {
			fmt.Fprintf(writer, "Failed to insert data into database: %v", err)
			return
		}
	}

}

func errors(boodschap string, err error, exitcode int) {
	fmt.Println(boodschap, err)
	os.Exit(exitcode)
}
