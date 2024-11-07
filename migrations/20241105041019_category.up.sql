-- Создание таблицы categories
CREATE TABLE if not exists categories (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL UNIQUE
);