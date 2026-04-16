-- name: MarkFeedFetch :exec
UPDATE feeds 
SET updated_at = $1,
last_fetched_at = $1
WHERE feeds.id = $2;
