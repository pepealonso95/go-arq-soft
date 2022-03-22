Primero que nada correr `go mod tidy`.
Esto va a traer todas las dependencias necesarias, es como correr npm install pero mejor ;) 

Para correr tenemos dos opciones:
- Correr: `go run main.go`


-  Compilar y correr: 
   1) `go build main.go`
   2) `./main` en mac y linux, `.\main.exe` en windows (no usen windows por favor )

Para probar los endpoints por consola (windows usa postman):
- GetBooks: `curl -X GET http://127.0.0.1:8080/api/v1/books`
- GetBook: `curl -X GET http://127.0.0.1:8080/api/v1/book/Hola`
- AddBook: `curl -X POST http://127.0.0.1:8080/api/v1/books -d '{"title":"Mi primer libro","author":"Profesor"}'`

Desafio:
Crear un endpoint que traiga todos los libros para un autor