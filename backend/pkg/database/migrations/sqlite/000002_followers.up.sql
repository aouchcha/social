CREATE TABLE IF NOT EXISTS followers (
	followed_id INTEGER NOT NULL,
	follower_id INTEGER NOT NULL,
	PRIMARY KEY (followed_id, follower_id),
	FOREIGN KEY (followed_id) REFERENCES users (user_id) ON DELETE CASCADE,
	FOREIGN KEY (follower_id) REFERENCES users (user_id) ON DELETE CASCADE
);