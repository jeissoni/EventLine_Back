package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jeissoni/EventLine/cmd/api/handlers/events"
	db "github.com/jeissoni/EventLine/internal/infrastructure/config/database"
	eventRepository "github.com/jeissoni/EventLine/internal/repositories/postgres/events"
	eventService "github.com/jeissoni/EventLine/internal/services/events"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// initialize fiber app
	app := fiber.New()

	contextdb, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	defer contextdb.Close()

	eventRepo := &eventRepository.Repository{
		Database: contextdb,
	}

	eventsSrv := &eventService.Service{
		Repository: eventRepo,
	}

	eventsHandler := &events.EventHandler{
		EventService: eventsSrv,
	}

	// define routes
	app.Post("/events", eventsHandler.CreateEvent)
	app.Get("/events/:id", eventsHandler.GetByID)
	app.Delete("/events/:id", eventsHandler.Delete)
	app.Get("/events", eventsHandler.GetAll)

	log.Fatalln(app.Listen(":3000"))

}
