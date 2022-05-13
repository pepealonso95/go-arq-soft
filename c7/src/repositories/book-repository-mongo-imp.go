package repositories

import (
	"c7/src/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collection = "books"
)

type bookRepoMongoImp struct {
	mongoClient *mongo.Client
	database    string
}

func NewBookMongoRepo(mongoClient *mongo.Client, database string) *bookRepoMongoImp {
	return &bookRepoMongoImp{
		mongoClient: mongoClient,
		database:    database,
	}
}

func (bookRepo *bookRepoMongoImp) AddBook(book models.Book) error {
	_, err := bookRepo.mongoClient.Database(bookRepo.database).Collection(collection).InsertOne(context.TODO(), book)
	return err
}

func (bookRepo *bookRepoMongoImp) GetBooks() ([]models.Book, error) {
	cursorBooks, err := bookRepo.mongoClient.Database(bookRepo.database).Collection(collection).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	var books []models.Book
	err = cursorBooks.All(context.TODO(), &books)

	return books, err
}

func (bookRepo *bookRepoMongoImp) GetBook(title string) (models.Book, error) {
	var book models.Book
	query := bson.M{"title": title}
	err := bookRepo.mongoClient.Database(bookRepo.database).
		Collection(collection).FindOne(context.TODO(), query).Decode(&book)
	return book, err
}
