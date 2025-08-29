-- 1581. Customer Who Visited But Did Not Make Any Transactions
-- Link: https://leetcode.com/problems/customer-who-visited-but-did-not-make-any-transactions
-- Database : MySQL

SELECT v.customer_id, COUNT(v.customer_id) count_no_trans
FROM Visits v
LEFT JOIN Transactions t
ON v.visit_id = t.visit_id 
WHERE t.transaction_id is null
GROUP BY v.customer_id