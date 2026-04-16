-- name: ListFeeds :many
SELECT feeds.name, feeds.url, users.name
FROM users
INNER JOIN feeds
on users.id = feeds.user_id;