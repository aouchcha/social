CREATE TABLE IF NOT EXISTS group_events (
	event_id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	group_id INTEGER NOT NULL,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	event_date DATE NOT NULL,
	event_time TIME NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
	FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE CASCADE
);