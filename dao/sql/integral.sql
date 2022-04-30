drop DATABASE if exists DBIntegralFlow_$appid;
create database DBIntegralFlow_$appid;
create table if not exists DBIntegralFlow_$appid.tbIntegralFlow_$i
(
    appid     VARCHAR(64)     NOT NULL DEFAULT '',
    type      VARCHAR(64)     NOT NULL DEFAULT '',
    id        VARCHAR(64)     NOT NULL DEFAULT '',
    oid       VARCHAR(64)     NOT NULL DEFAULT '',
    opt       INT             NOT NULL DEFAULT 0,
    integral  BIGINT          NOT NULL DEFAULT 0,
    time      VARCHAR(64)     NOT NULL DEFAULT '',
    timestamp BIGINT UNSIGNED NOT NULL DEFAULT 0,
    rollback  BOOL            NOT NULL DEFAULT false,
    primary key (appid, type, id, oid),
    index (appid, type, id, timestamp)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

DROP DATABASE IF EXISTS DBIntegral_$appid;
CREATE DATABASE DBIntegral_$appid;
CREATE TABLE IF NOT EXISTS DBIntegral_$appid.tbIntegral_$i
(
    appid    VARCHAR(64) NOT NULL DEFAULT '',
    type     VARCHAR(64) NOT NULL DEFAULT '',
    id       VARCHAR(64) NOT NULL DEFAULT '',
    integral BIGINT      NOT NULL DEFAULT 0,
    PRIMARY KEY (appid, type, id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;