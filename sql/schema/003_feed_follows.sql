-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL,
    feed_id UUID NOT NULL,
    FOREIGN KEY (USER_ID) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (FEED_ID) REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE (USER_ID, FEED_ID)
);

-- +goose Down
DROP TABLE feed_follows;