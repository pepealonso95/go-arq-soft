package main

import (
	"time"

	"c3/src/configs"
	"c3/src/controllers"
	"c3/src/middlewares"
	"c3/src/models"
	"c3/src/repositories"
	"c3/src/routes"

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

	// Repositories
	// Defino mi instancia del repositorio para inyectarla a los controladores a utilizar
	// Tengo que sacar la direccion de memoria a mano para pasarlo a las funciones
	// que utilizan mi instancia de repo, al ser metodo y no funcion se pierde el syntax sugar
	repo := &repositories.BookRepo{}

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

func loadBooks(repo *repositories.BookRepo) {
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
