-- 1148. Article Views I
-- Link: https://leetcode.com/problems/article-views-i
-- Database : MySQL

SELECT DISTINCT(author_id) id
FROM Views
WHERE author_id = viewer_id
ORDER BY id ASC
