-- Database example
-- College(cName,state,enrollment)
-- Student(sID,sName,GPA,sizeHS)
-- Apply(sID,cName,major,decision)


-- # Example 1
select
  sID,
  cName,
  GPA
from Student; -- FROM Statement

-- # Example 2
select
  cName,
  state,
  enrollment
from College; -- FROM Statement

-- # Example 3
select
  cName,
  major,
  decision
from Apply; -- FROM Statement




