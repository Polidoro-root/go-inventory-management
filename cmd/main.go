package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Polidoro-root/go-inventory-management/configs"
	"github.com/Polidoro-root/go-inventory-management/internal/infra/database"
	"github.com/Polidoro-root/go-inventory-management/internal/infra/web"
	"github.com/Polidoro-root/go-inventory-management/internal/infra/web/webserver"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	configs := configs.LoadConfig()

	db, err := sql.Open(
		configs.DBDriver,
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s %s",
			configs.DBUser,
			configs.DBPassword,
			configs.DBHost,
			configs.DBPort,
			configs.DBName,
			configs.DBOptions,
		),
	)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	userRepository := database.NewUserRepository(db)

	webUserHandler := web.NewWebUserHandler(userRepository)

	server := webserver.NewWebServer(configs.WebServerPort)

	server.AddHandler("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// webUserHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	server.AddHandler("/users/signin", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			webUserHandler.SignIn(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	server.Start()
}
