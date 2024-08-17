-- Database example
-- College(cName,state,enrollment)
-- Student(sID,sName,GPA,sizeHS)
-- Apply(sID,cName,major,decision)

-- # Example 1
select
  sID,
  cName,
  GPA
from Student
where GPA > 5; -- Where statement

-- # Example 2
select
  sID,
  cName
from Student
where GPA > 2 AND sizeHS > 1000; -- Use logical operator 'AND'

