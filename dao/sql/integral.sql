drop DATABASE if exists DBIntegralFlow_$appid;
create database DBIntegralFlow_$appid;
create table if not exists DBIntegralFlow_$appid.tbIntegralFlow_$i
(
    appid     VARCHAR(64)     NOT NULL DEFAULT '',
    type      VARCHAR(64)     NOT NULL DEFAULT '',
    id        VARCHAR(64)     NOT NULL DEFAULT '',
    oid       VARCHAR(64)     NOT NULL DEFAULT '',
    opt       INT             not null default 0,
    integral  BIGINT          not null default 0,
    time      VARCHAR(64)     NOT NULL DEFAULT '',
    timestamp BIGINT UNSIGNED NOT NULL DEFAULT 0,
    rollback  bool            not null default false,
    primary key (appid, type, id, oid),
    index (appid, type, id, timestamp)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

drop database if exists DBIntegral_$appid;
create database DBIntegral_$appid;
create table if not exists DBIntegral_$appid.tbIntegral_$i
(
    appid    VARCHAR(64) NOT NULL DEFAULT '',
    type     VARCHAR(64) NOT NULL DEFAULT '',
    id       VARCHAR(64) NOT NULL DEFAULT '',
    integral BIGINT      not null default 0,
    primary key (appid, type, id)
) engine = innodb
  default charset = utf8;