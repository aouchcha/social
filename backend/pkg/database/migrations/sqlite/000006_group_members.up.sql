CREATE TABLE IF NOT EXISTS group_members (
	group_id INTEGER NOT NULL,
	member_id INTEGER NOT NULL,
	PRIMARY KEY (group_id, member_id),
	FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE CASCADE,
	FOREIGN KEY (member_id) REFERENCES users (user_id) ON DELETE CASCADE
);