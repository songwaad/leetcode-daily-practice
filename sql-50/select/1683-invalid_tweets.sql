-- 1683. Invalid Tweets
-- Link: https://leetcode.com/problems/invalid-tweets
-- Database : MySQL

SELECT tweet_id
FROM Tweets 
WHERE length(content) > 15