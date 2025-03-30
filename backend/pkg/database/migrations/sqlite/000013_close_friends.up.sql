CREATE TABLE IF NOT EXISTS close_friends (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL,
    follower_id INTEGER NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts (post_id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES followers (follower_id) ON DELETE CASCADE
);