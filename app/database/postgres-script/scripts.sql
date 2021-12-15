CREATE TABLE
IF NOT EXISTS users
(
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL CHECK (name <> ''),
    PRIMARY KEY(id)
);

CREATE TABLE
IF NOT EXISTS posts
(
    id INT GENERATED ALWAYS AS IDENTITY,
    user_id INT,
    body text,
    PRIMARY KEY(id),
    CONSTRAINT fk_user
       FOREIGN KEY(user_id)
	   REFERENCES users(id)
);



INSERT INTO users(name) VALUES('foo');
INSERT INTO users(name) VALUES('bar');

INSERT INTO posts(user_id, body) VALUES(1, 'Hello World');
INSERT INTO posts(user_id, body) VALUES(2, 'Hi');
