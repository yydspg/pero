use pero;

-- table info
drop table if exists `items`;
create table `items` (
                         id bigint unsigned not null primary key AUTO_INCREMENT,
                         item_id bigint unsigned not null default 0,
                         service_id bigint unsigned not null default 0,
                         short_url varchar(20) not null  comment '短链',
                         dest_url text not null comment '长链',
                         is_valid smallint not null default 0 comment '状态,0.ok,1.error',
                         version int not null default 0,
                         create_at timestamp not null,
                         update_at timestamp not null
);
-- service info
drop table if exists `services`;
create table `services` (
                            id bigint unsigned not null primary key AUTO_INCREMENT,
                            service_id bigint unsigned not null default 0,
                            service_name varchar(20) not null default '',
                            tag varchar(20) not  null  comment '服务标记',
                            num int unsigned not null default 0 comment '服务Item数量',
                            status smallint not null default 0 comment '状态 0.正常,1.失效',
                            create_at timestamp not null ,
                            update_at timestamp not null
);
