CREATE TABLE IF NOT EXISTS sessions (
	session_id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	token TEXT UNIQUE NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
);