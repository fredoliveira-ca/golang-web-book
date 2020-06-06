create table book (
	id text primary key,
	title character varying(120),
	author character varying(120),
	price decimal
)

INSERT INTO
	book(id, title, author, price)
VALUES
	(
		'b6fea6da-aeb7-4ff6-be0f-882c56c574ea',
		'Refactoring',
		'Martin Fowler',
		4335.21
	);

INSERT INTO
	book(id, title, author, price)
VALUES
	(
		'22021344-4a1f-4979-bacf-20f30353b751',
		'Domain-driven design',
		'Eric Evans',
		61.2
	);