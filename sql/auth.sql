# 新建库
CREATE DATABASE IF NOT EXISTS auth;
use auth;

# 用户表
CREATE TABLE IF NOT EXISTS user (
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name        VARCHAR(50) NOT NULL,
    email       VARCHAR(50) NOT NULL,
    passwd      CHAR(32) NOT NULL,
    host        VARCHAR(50) NOT NULL,
    is_vertify  BOOLEAN NOT NULL,
    is_ban      BOOLEAN NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# 注册初始用户
INSERT user(name, email, passwd, host, is_vertify, is_ban) VALUES('赤い彗星', 'leonardocaesarz@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', '127.0.0.1', 1, 0)
