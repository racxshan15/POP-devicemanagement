package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HTMLpagina)         //Zorgt ervoor dat het naar de HTMLpagina gaat.
	err := http.ListenAndServe(":8080", nil) //Hier staat dat het moet draaien op port 8080.
	if err != nil {
		errors("Kon niet verbinden met de port", err, 0) //Nog niet zeker of dit nut heeft.
	}
}

func HTMLpagina(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "C:\\Users\\racxs\\OneDrive\\Documenten\\Infrastructure\\Semester 2\\Pop\\web.html") //Dit is de html bestand.
}

func errors(boodschap string, err error, exitcode int) {
	fmt.Println(boodschap, err)
	os.Exit(exitcode)
}
