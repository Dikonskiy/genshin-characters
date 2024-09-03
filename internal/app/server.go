package app

import (
	"genshin/internal/config"
	"genshin/internal/database"
	"genshin/internal/handlers"
	"genshin/internal/parser"
	"genshin/internal/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) StartServer() {
	cnfg, err := config.NewConfig(".env")
	if err != nil {
		return
	}
	db, err := database.NewMysqlConn(cnfg)
	parser := parser.NewParser()
	repo := repository.NewRepository(db.GetDB())
	hand := handlers.NewHandler(repo, parser)

	r := mux.NewRouter()

	server := &http.Server{
		Addr:         ":" + cnfg.GetPort(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
	}

	quit := make(chan os.Signal, 1)

	go shutdown(quit)

	r.HandleFunc("/characters", hand.InsertCharacterHandler).Methods("POST")

	log.Println("Listening on port", cnfg.GetPort(), "...")
	log.Fatal(server.ListenAndServe())
}

func shutdown(quit chan os.Signal) {
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	log.Println("caught signal", s.String())
	os.Exit(0)
}
