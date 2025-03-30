CREATE TABLE IF NOT EXISTS group_invitations (
	group_id INTEGER NOT NULL,
	invited INTEGER NOT NULL,
	invited_by INTEGER NOT NULL,
	status TEXT DEFAULT 'p' CHECK(status IN ('p', 'a', 'r')),
	PRIMARY KEY (group_id, invited),
	FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE CASCADE,
	FOREIGN KEY (invited) REFERENCES users (user_id) ON DELETE CASCADE,
	FOREIGN KEY (invited_by) REFERENCES users (user_id) ON DELETE CASCADE
);