CREATE TABLE IF NOT EXISTS users (
	user_id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	date_of_birth DATETIME NOT NULL,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	avatar_url TEXT,
	password TEXT NOT NULL,
	about TEXT,
	public INTEGER DEFAULT 1 CHECK (public IN (0, 1))
);