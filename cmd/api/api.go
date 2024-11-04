package api

import (
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/oseayemenre/go_crud_scratch/internal/routes"
	"github.com/oseayemenre/go_crud_scratch/internal/sql/database"
	"github.com/oseayemenre/go_crud_scratch/internal/types"
)

func BootstrapServer() {
	godotenv.Load()

	port := os.Getenv("PORT")
	db_url := os.Getenv("DB_URL")

	if port == "" {
		port = "8080"
	}

	if db_url == "" {
		log.Fatal("Unable to connect to DB")
	}


	conn, err := sql.Open("postgres", db_url)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	apiCfg := &types.ApiConfig{
		DB: database.New(conn),
	}

	db := &routes.DB{
		ApiConfig: apiCfg,
	}

	r := chi.NewRouter()

	db.HandleRoutes(r)

	log.Printf("Server is running on PORT: %v...", port)
	if err := http.ListenAndServe(":" + port, r); err != nil {
		log.Fatalf("error: %v", err)
	}
}