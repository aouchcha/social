CREATE TABLE IF NOT EXISTS group_requests (
	group_id INTEGER NOT NULL,
	requester_id INTEGER NOT NULL,
	status TEXT DEFAULT 'p' CHECK(status IN ('p', 'a', 'r')),
	PRIMARY KEY (group_id, requester_id),
	FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE CASCADE,
	FOREIGN KEY (requester_id) REFERENCES users (user_id) ON DELETE CASCADE
);