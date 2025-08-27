-- 595. Big Countries
-- Link: https://leetcode.com/problems/big-countries
-- Database : MySQL

SELECT name, population, area
FROM World
WHERE area > 3000000 || population > 25000000