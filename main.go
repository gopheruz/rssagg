package main

import (
	"database/sql"
	"gihub/com/nurmuhammaddeveloper/rssag/internal/databse"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	Db *databse.Queries
}

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("filed to get port from envitoment")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("filed to get databse url from enviroment")
	}
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("filed to connect databse")
	}
	apiCfg := apiConfig{
		Db: databse.New(conn),
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"LINK"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()

	v1Router.HandleFunc("/", responseWithJsonHandler)
	v1Router.Post("/users", apiCfg.handlerCreateuser)
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("startting server at http://localhost:%s", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Filed to starting server \n", err.Error())
	}

}
