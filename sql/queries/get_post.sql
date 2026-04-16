-- name: GetPost :many
SELECT posts.*, feeds.name as feed_name
FROM posts
INNER JOIN feed_follows
ON posts.feed_id = feed_follows.feed_id
INNER JOIN feeds
on posts.feed_id = feeds.id
WHERE feed_follows.user_id = $2
ORDER BY posts.published_at desc
limit $1;