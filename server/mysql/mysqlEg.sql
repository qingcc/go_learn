/*
编写一个 SQL 查询来实现分数排名。如果两个分数相同，则两个分数排名（Rank）相同。请注意，平分后的下一个名次应该是下一个连续的整数值。换句话说，名次之间不应该有“间隔”。

+----+-------+
| Id | Score |
+----+-------+
| 1  | 3.50  |
| 2  | 3.65  |
| 3  | 4.00  |
| 4  | 3.85  |
| 5  | 4.00  |
| 6  | 3.65  |
+----+-------+
例如，根据上述给定的 Scores 表，你的查询应该返回（按分数从高到低排列）：

+-------+------+
| Score | Rank |
+-------+------+
| 4.00  | 1    |
| 4.00  | 1    |
| 3.85  | 2    |
| 3.65  | 3    |
| 3.65  | 3    |
| 3.50  | 4    |
+-------+------+

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rank-scores
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

select score, (select count(distinct score) from scores where score >= s.score) AS Rank from scores s order by score desc;


/*
编写一个 SQL 查询，查找所有至少连续出现三次的数字。

+----+-----+
| Id | Num |
+----+-----+
| 1  |  1  |
| 2  |  1  |
| 3  |  1  |
| 4  |  2  |
| 5  |  1  |
| 6  |  2  |
| 7  |  2  |
+----+-----+
例如，给定上面的 Logs 表， 1 是唯一连续出现至少三次的数字。

+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+

-- todo @prev, @count 为用户变量，  (select @prev := null,@count := null) as t   该段sql是用来初始化用户变量的
-- todo case when then , 如果当前的数字和之前的相等 num == @priv， @count + 1 ， 否则 @priv = num， @count = 1
-- todo  select Num,
-- todo    case
-- todo      when @prev = Num then @count := @count + 1
-- todo      when (@prev := Num) is not null then @count := 1
-- todo    end as CNT
-- todo  from Logs, (select @prev := null,@count := null) as t
-- todo 该段sql输出为：
+----+-----+
| Num| CNT |
+----+-----+
| 1  |  1  |
| 1  |  2  |
| 1  |  3  |
| 2  |  1  |
| 1  |  1  |
| 2  |  1  |
| 2  |  2  |
+----+-----+
*/
select distinct Num as ConsecutiveNums
from (
  select Num,
    case
      when @prev = Num then @count := @count + 1
      when (@prev := Num) is not null then @count := 1
    end as CNT
  from Logs, (select @prev := null,@count := null) as t
) as temp
where temp.CNT >= 3


/*
Employee 表包含所有员工信息，每个员工有其对应的 Id, salary 和 department Id。

+----+-------+--------+--------------+
| Id | Name  | Salary | DepartmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 70000  | 1            |
| 2  | Henry | 80000  | 2            |
| 3  | Sam   | 60000  | 2            |
| 4  | Max   | 90000  | 1            |
+----+-------+--------+--------------+
Department 表包含公司所有部门的信息。

+----+----------+
| Id | Name     |
+----+----------+
| 1  | IT       |
| 2  | Sales    |
+----+----------+
编写一个 SQL 查询，找出每个部门工资最高的员工。例如，根据上述给定的表格，Max 在 IT 部门有最高工资，Henry 在 Sales 部门有最高工资。

+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Max      | 90000  |
| Sales      | Henry    | 80000  |
+------------+----------+--------+

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/department-highest-salary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

-- todo 自己完成
select temp.Name as Department, e.Name as Employee, temp.Salary from Employee as e,
(select max(Salary) as Salary, DepartmentId, d.Name as name from Employee as e, Department as d
 where e.DepartmentId = d.Id group by DepartmentId)
as  temp
where e.Salary = temp.Salary and e.DepartmentId = temp.DepartmentId





