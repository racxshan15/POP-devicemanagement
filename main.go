package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", HTMLpagina)         //Zorgt ervoor dat het naar de HTMLpagina gaat.
	err := http.ListenAndServe(":8080", nil) //Hier staat dat het moet draaien op port 8080.
	if err != nil {
		errors("Kon niet verbinden met de port", err, 0) //Nog niet zeker of dit nut heeft.
	}

	// Maak een databaseverbinding
	db, err := sql.Open("mysql", "root:v01LnMA1XG*T@tcp(localhost:3306)/apparaattoekenningen")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	db.Query("INSERT INTO `apparaattoekenningen`.`toekenningen`(`voornaam`,`achternaam`,`apparaat`,`toekenning_datum`,`geldig_tot`) VALUE('hoi','go','laptop',NOW(),NOW());")
}

func HTMLpagina(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "C:\\Users\\racxs\\OneDrive\\Documenten\\Infrastructure\\Semester 2\\Pop\\web.html") //Dit is de html bestand.
}

func errors(boodschap string, err error, exitcode int) {
	fmt.Println(boodschap, err)
	os.Exit(exitcode)
}
