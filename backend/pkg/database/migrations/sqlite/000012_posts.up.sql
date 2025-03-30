CREATE TABLE IF NOT EXISTS posts (
    post_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    image_url TEXT,
    group_id INTEGER,
    privacy INTEGER DEFAULT 'pb' CHECK(privacy IN ('pr', 'pb','ap')), --op=almost_public pr=private, pb=public
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE CASCADE
);