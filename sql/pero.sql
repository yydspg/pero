create database 'pero';
use pero;

-- table info
drop table if exists 'item';
create table 'item' (
    id bigint unsigned not null primary key AUTO_INCREMENT,
    item_id bigint unsigned not null default 0,
    service_id bigint unsigned not null default 0,
    short_url varchar(20) not null default '' comment '短链',
    dest_url text not null default ''comment '长链',
    is_valid smallint not null default 0 comment '状态,0.ok,1.error',
    version int not null default 0,
    create_at timestamp not null default CURRENT_TIMESTAMP,
    update_at timestamp not null default CURRENT_TIMESTAMP
);
-- service info
drop table if exists 'service';
create table 'service' (
    id bigint unsigned not null primary key AUTO_INCREMENT,
    service_id bigint unsigned not null default 0,
    service_name varchar(20) not null default '',
    tag varchar(20) not  null default '' comment '服务标记',
    status smallint not null default 0 comment '状态 0.正常,1.失效',
    create_at timestamp not null default CURRENT_TIMESTAMP,
    update_at timestamp not null default CURRENT_TIMESTAMP
)
