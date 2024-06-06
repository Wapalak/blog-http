-- Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS users(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

-- Создание таблицы блогов
CREATE TABLE IF NOT EXISTS blog (
    id SERIAL PRIMARY KEY,
    author VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    text TEXT NOT NULL,
    votes INTEGER DEFAULT 0
);
