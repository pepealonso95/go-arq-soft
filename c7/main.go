package main

import (
	"context"
	"time"

	"c7/src/configs"
	"c7/src/controllers"
	"c7/src/controllers/repointerfaces"
	"c7/src/datasources"
	"c7/src/middlewares"
	"c7/src/models"
	"c7/src/repositories"
	"c7/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	// Aqui defino middlewares que quiera para mi app, solo uso un logger por ahora
	middlewares.FiberMiddleware(app) // Register Fiber's middleware for app.

	mongoClient, err := datasources.NewMongoDataSource("mongodb://docker:mongopw@localhost:55000")

	if err != nil {
		panic(err)
	}

	defer mongoClient.Disconnect(context.TODO())

	repo := repositories.NewBookMongoRepo(mongoClient, "booksDB")

	// Esto es para cargar los libros en memoria, al usar una db no sera necesario
	loadBooks(repo)

	// Creo una instancia de mis controladores con mi instancia de repo
	controller := controllers.NewBookController(repo)

	// Routes.
	// Aqui defino cuales van a ser las rutas accesibles
	routes.PublicRoutes(app, controller) // Register a public routes for app.

	// Aqui inicializamos el servidor en el puerto 8080
	app.Listen(":8080")
}

func loadBooks(repo repointerfaces.BookRepo) {
	repo.AddBook(models.Book{
		Title:     "Nadie lee",
		Author:    "Alumno",
		CreatedAt: time.Now(),
	})
	repo.AddBook(models.Book{
		Title:     "Pasenme el obli anterior",
		Author:    "Alumno",
		CreatedAt: time.Now().Add(-time.Hour),
	})
	repo.AddBook(models.Book{
		Title:     "Hola",
		Author:    "Profesor",
		CreatedAt: time.Now().Add(-time.Hour * 2),
	})
	repo.AddBook(models.Book{
		Title:     "Se van todos a examen",
		Author:    "Profesor",
		CreatedAt: time.Now().Add(-time.Hour * 3),
	})
	repo.AddBook(models.Book{
		Title:     "Que es estudiar?",
		Author:    "Alumno de diseno",
		CreatedAt: time.Now().Add(-time.Hour * 4),
	})
}
