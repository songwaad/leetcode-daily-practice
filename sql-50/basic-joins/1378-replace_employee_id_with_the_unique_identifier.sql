-- 1378. Replace Employee ID with The Unique Identifier
-- Link: https://leetcode.com/problems/replace-employee-id-with-the-unique-identifier
-- Database : MySQL

SELECT empUNI.unique_id, emp.name
FROM Employees emp
LEFT JOIN EmployeeUNI empUNI
ON emp.id = empUNI.id