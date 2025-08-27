-- 584. Find Customer Referee
-- Link: https://leetcode.com/problems/find-customer-referee
-- Database : MySQL

SELECT name
FROM Customer
WHERE referee_id != 2 || referee_id is null;