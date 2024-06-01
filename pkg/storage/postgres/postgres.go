package postgres

import (
	"GoNews/pkg/storage"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

// Конструктор
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Получение всех постов
func (s *Storage) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `SELECT id,author_id,title,content,created FROM posts;
		`)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for rows.Next() {
		var post storage.Post
		err = rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, rows.Err()
}

// Добавление поста
func (s *Storage) AddPost(t storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
	INSERT INTO posts(author_id,title,content,created) VALUES($1,$2,$3,$4);
		`,
		t.AuthorID,
		t.Title,
		t.Content,
		t.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil

}

// Обновление поста
func (s *Storage) UpdatePost(t storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
	UPDATE posts SET author_id=$1,title=$2,content=$3,created=$4 WHERE id=$5;
		`,
		t.AuthorID,
		t.Title,
		t.Content,
		t.CreatedAt,
		t.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// Удаление поста
func (s *Storage) DeletePost(t storage.Post) error {
	_, err := s.db.Exec(context.Background(),
		`DELETE FROM posts WHERE id=$1;`,
		t.ID)
	if err != nil {
		return err
	}
	return nil
}
