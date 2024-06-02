package mongo

import (
	"GoNews/pkg/storage"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName   = "gonews3" // имя учебной БД
	collectionName = "posts"   // имя коллекции в учебной БД
)

type Storage struct {
	db *mongo.Client
}

// Конструктор для mongo
func New(constr string) (*Storage, error) {
	mongoOpts := options.Client().ApplyURI(constr)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: client,
	}
	return &s, err
}

// Вывод всех постов
func (s *Storage) Posts() ([]storage.Post, error) {
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var posts []storage.Post
	for cur.Next(context.Background()) {
		var p storage.Post
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, cur.Err()
}

// Добавление поста
func (s *Storage) AddPost(p storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

// Обновление контента поста
func (s *Storage) UpdatePost(p storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.D{{"ID", p.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"Content", p.Content},
		}},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// Удаление поста
func (s *Storage) DeletePost(p storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.D{{`ID`, p.ID}}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil

}
