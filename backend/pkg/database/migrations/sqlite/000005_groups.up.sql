CREATE TABLE IF NOT EXISTS groups (
	group_id INTEGER PRIMARY KEY AUTOINCREMENT,
	owner_id INTEGER NOT NULL,
	name TEXT UNIQUE NOT NULL,
	description TEXT NOT NULL,
	FOREIGN KEY (owner_id) REFERENCES users (user_id) ON DELETE CASCADE
);
