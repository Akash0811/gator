-- name: GetFeedFollowsForUser :many
SELECT feeds.name as feed_name
FROM feed_follows
INNER JOIN feeds
on feed_follows.feed_id = feeds.id
INNER JOIN users
on feed_follows.user_id = users.id
where users.name = $1;