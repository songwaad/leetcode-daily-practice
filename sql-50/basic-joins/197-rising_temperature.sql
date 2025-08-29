-- 197. Rising Temperature
-- Link: https://leetcode.com/problems/rising-temperature
-- Database : MySQL

SELECT today.id
FROM Weather today
JOIN Weather yesterday
WHERE DATEDIFF(today.recordDate, yesterday.recordDate) = 1
    AND (today.temperature > yesterday.temperature)