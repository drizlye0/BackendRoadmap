-- Database example
-- College(cName,state,enrollment)
-- Student(sID,sName,GPA,sizeHS)
-- Apply(sID,cName,major,decision)

-- Select sID, cName from relation STUDENT where GPA > 3.6
select
  sID,
  cName,
  GPA -- Is not neecessary include
from Student
where GPA > 3.6;

-- Select cName, major from relations Student and Apply
-- where Student.sID and Apply.sID is the same
select
  cName,
  major
from Student, Apply
where Student.sID = Apply.sID;

-- but, this query also responds with duplicate data
-- adds a *distinct* keyword on the query for fix this

select distinct
   sName,
   major
from Student, Apply
where Student.sID = Apply.sID;

-- Select name of college from relations College and Apply
-- where College.cName and Apply.cName is the same and the
-- enrollment is greather than 2000 and major is CS

select
  cName
from College, Apply
where College.cName = Apply.cName
  and enrollment > 2000 and major = "CS";

-- this query responds with an error, the error corresponds
-- at select cName, SQL find two ambiguous column names, and 
-- finds no difference with between College table and Apply table

-- for fix this error, you must specify the table for select
-- the column name
-- select
--    *College.cName*
-- .....
