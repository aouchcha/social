CREATE TABLE IF NOT EXISTS messages (
	message_id INTEGER PRIMARY KEY AUTOINCREMENT,
	sender_id INTEGER NOT NULL,
	receiver_id INTEGER NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (sender_id) REFERENCES users (user_id) ON DELETE CASCADE,
	FOREIGN KEY (receiver_id) REFERENCES users (user_id) ON DELETE CASCADE
);