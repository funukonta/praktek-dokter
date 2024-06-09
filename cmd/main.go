package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/funukonta/praktek-dokter/pkg"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mux := http.NewServeMux()

	db := pkg.ConnectPostgre()
	fmt.Println(db)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	})

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.ListenAndServe(port, mux)
}

//
// docker run --name praktek-app -p 5432:5432 -e POSTGRES_PASSWORD=postgre -d postgres
