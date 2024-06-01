DROP TABLE IF EXISTS authors,posts;

CREATE TABLE authors(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL
);

CREATE TABLE posts(
	id SERIAL PRIMARY KEY,
	author_id INTEGER REFERENCES authors(id) NOT NULL,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created BIGINT NOT NULL
);

INSERT INTO authors(id,name) VALUES(0,'Антон');
INSERT INTO authors(id,name) VALUES(1,'Дмитрий');
INSERT INTO authors(id,name) VALUES(2,'Игорь');

INSERT INTO posts(author_id,title,content,created) VALUES(0,'Заголовок 1','Статья 1',0);
INSERT INTO posts(author_id,title,content,created) VALUES(1,'Заголовок 2','Статья 2',0);
INSERT INTO posts(author_id,title,content,created) VALUES(2,'Заголовок 3','Статья 3',0);

