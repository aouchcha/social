CREATE TABLE IF NOT EXISTS likes_comment (
	like_id INTEGER PRIMARY KEY AUTOINCREMENT,
	comment_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	is_like INTEGER CHECK (is_like IN (1)),
	FOREIGN KEY (comment_id) REFERENCES comments (comment_id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
);