-- 1069. Product Sales Analysis I
-- Link: https://leetcode.com/problems/product-sales-analysis-i
-- Database : MySQL

SELECT p.product_name, s.year, s.price
FROM Sales s
LEFT JOIN Product p
ON s.product_id = p.product_id