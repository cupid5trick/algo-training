-- create table sale_info (
-- Id bigint unsigned not null primary key auto_increment,
-- OrderNo bigint not null default 0 comment '销售单号',
-- ProductId int unsigned not null default 0 comment '商品ID',
-- Num smallint unsigned not null default 1 comment '销售数量',
-- Price numeric(8, 2) not null default 0 comment '单价',
-- SaleTime datetime comment '销售时间',

-- key IdxOrderNo(OrderNo),
-- key IdxProductId(ProductId),
-- key IdxSaleTime(SaleTime)
-- )engine=innodb charset=utf8 comment='商品销售详情表';

-- create table product_info (
-- Id int unsigned not null primary key auto_increment comment '商品ID',
-- ProductName varchar(128) not null default '' comment '商品名称',

-- key IdxProductName(ProductName(32))
-- )engine=innodb charset=utf8 comment='商品信息表';
-- 插入少量数据

-- insert into product_info(Id, ProductName)values(1, '足球'), (2, '篮球'), (3, '网球'), (4, '乒乓球'), (5, '羽毛球');
-- insert into sale_info(OrderNo, ProductId, Num, Price, SaleTime)values
-- ('5000123100', 1, 1, 128, '2024-02-14 19:39:01'),
-- ('5000123101', 2, 5, 298, '2024-02-14 20:43:01'),
-- ('5000123102', 5, 2, 8, '2024-02-15 19:39:01'),
-- ('5000123103', 4, 1, 6.5, '2024-02-04 19:11:01'),
-- ('5000123104', 2, 1, 298, '2024-03-07 19:39:01'),
-- ('5000123105', 3, 1, 10, '2024-03-14 09:39:01');

-- 问题：

-- 写一条update语句，将数据表sale_info中 主键Id 值为 1 的记录，销售数量 Num 字段更新为 3
-- 通过一条SQL语句，统计2024年2月份的总销售额
-- 通过一条SQL语句，统计2024年2月份销售数量最多的的前3个商品ID以及名称

-- 1 innodb next-key
A:5 B:6(commit) A: select 6;

select ... for update; -- id=1 (1,+inf)
update sale_info set Num=3 where Id = 1;

--2 
select sum(Price*Num) from sale_info where SaleTime >= '2024-02-01' and SaleTime < '2024-03-01';

--3 
select Id, ProductName (select Id, ProductName, sum(Num) volumn from sale_info inner join product_info on Id where SaleTime >= '2024-02-01' and SaleTime < '2024-03-01' group by Id order by volumn limit 3);
