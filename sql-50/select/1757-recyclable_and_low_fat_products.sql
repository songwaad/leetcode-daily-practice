-- 1757. Recyclable and Low Fat Products
-- Link: https://leetcode.com/problems/recyclable-and-low-fat-products
-- Database : MySQL

SELECT product_id
FROM Products
WHERE low_fats = 'Y' && recyclable = 'Y';