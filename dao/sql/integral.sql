CREATE DATABASE if not exists DBIntegralFlow_$i;
create table if not exists DBIntegralFlow_$i.tbIntegralFlow_$j
(
    id        VARCHAR(64)     NOT NULL DEFAULT '',
    oid       VARCHAR(64)     NOT NULL DEFAULT '',
    appid     VARCHAR(64)     NOT NULL DEFAULT '',
    type      VARCHAR(64)     NOT NULL DEFAULT '',
    opt       INT             not null default 0,
    integral  BIGINT UNSIGNED not null default 0,
    timestamp BIGINT UNSIGNED NOT NULL DEFAULT 0,
    time      VARCHAR(64)     NOT NULL DEFAULT '',
    desc      VARCHAR(256)    not null default '',
    primary key (id, timestamp, oid)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

create database if not exists DBIntegral_$i;
create table if not exists DBIntegral_$i.tbIntegral_$j
(
    id    VARCHAR(64)  NOT NULL DEFAULT '',
    appid VARCHAR(64)  NOT NULL DEFAULT '',
    type  VARCHAR(64)  NOT NULL DEFAULT '',
    desc  VARCHAR(256) not null default '',
    primary key (type, id)
) engine = innodb
  default charset = utf8;