CREATE DATABASE if not exists DBIntegralFlow_$i;
create table if not exists DBIntegralFlow_$i.tbIntegralFlow_$j
(
    appid     VARCHAR(64)     NOT NULL DEFAULT '',
    type      VARCHAR(64)     NOT NULL DEFAULT '',
    id        VARCHAR(64)     NOT NULL DEFAULT '',
    oid       VARCHAR(64)     NOT NULL DEFAULT '',
    opt       INT             not null default 0,
    integral  BIGINT UNSIGNED not null default 0,
    time      VARCHAR(64)     NOT NULL DEFAULT '',
    desc      VARCHAR(256)    not null default '',
    timestamp BIGINT UNSIGNED NOT NULL DEFAULT 0,
    rollback  bool            not null default false,
    primary key (appid, type, id, oid)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

create database if not exists DBIntegral_$i;
create table if not exists DBIntegral_$i.tbIntegral_$j
(
    appid    VARCHAR(64)     NOT NULL DEFAULT '',
    type     VARCHAR(64)     NOT NULL DEFAULT '',
    id       VARCHAR(64)     NOT NULL DEFAULT '',
    integral BIGINT UNSIGNED not null default 0,
    desc     VARCHAR(256)    not null default '',
    primary key (appid, type, id)
) engine = innodb
  default charset = utf8;