CREATE VIEW IF NOT EXISTS all_posts_with_details AS
SELECT 
    p.post_id,
    u.user_id AS author_id,
	p.group_id,
    u.username AS author_username,
    p.content,
	p.image_url,
	p.privacy,
    COUNT(DISTINCT pr.like_id) AS like_count,
    COUNT(DISTINCT c.comment_id) AS comment_count,
    p.created_at
FROM 
    posts p
LEFT JOIN users u ON p.user_id = u.user_id
LEFT JOIN likes_post  pr ON p.post_id = pr.post_id AND pr.is_like = 1
LEFT JOIN comments c ON p.post_id = c.post_id
GROUP BY 
    p.post_id
ORDER BY 
    p.created_at DESC;