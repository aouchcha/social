CREATE TABLE IF NOT EXISTS event_interest (
    event_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    status TEXT CHECK(status IN ('g', 'n')), -- g=goingm n=not_going
    PRIMARY KEY (event_id, user_id),
    FOREIGN KEY (event_id) REFERENCES group_events (event_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
);