-- 创建额外数据库
CREATE DATABASE IF NOT EXISTS analytics_db;

-- 创建只读用户
CREATE USER 'reader'@'%' IDENTIFIED BY 'read_only_pass';
GRANT SELECT ON app_db.* TO 'reader'@'%';

-- 创建示例表
USE app_db;
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);