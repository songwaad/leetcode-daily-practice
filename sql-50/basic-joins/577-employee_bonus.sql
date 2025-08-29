-- 577. Employee Bonus
-- Link: https://leetcode.com/problems/rising-temperature
-- Database : MySQL

SELECT e.name, b.bonus
FROM Employee e
LEFt JOIN Bonus b
ON e.empId = b.empId
WHERE b.bonus < 1000 || b.bonus is null