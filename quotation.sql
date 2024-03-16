-- in postgres as root/postgres user ---------------------------
DROP DATABASE IF EXISTS quotebox;
CREATE DATABASE quotebox;

DROP OWNED BY quotebox CASCADE;
DROP USER IF EXISTS quotebox;
CREATE USER quotebox WITH PASSWORD 'tadeo2002';
-- ------------------------------------------------------------------


-- login to quotebox as postgres-----------------------------------
GRANT ALL PRIVILEGES ON DATABASE quotebox TO quotebox;
GRANT ALL ON SCHEMA public TO quotebox;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO quotebox;
-- ---------------------------------------------------


-- login to database as quotebox user ------------------------
DROP TABLE IF EXISTS quotations;
CREATE TABLE quotations(
    quotation_id SERIAL PRIMARY KEY,
    insertion_date date NOT NULL DEFAULT NOW(),
    author_name text NOT NULL,
    category text NOT NULL,
    quote text NOT NULL
);

-- insert a quote for testing purposes
INSERT INTO quotations(author_name, category, quote)
VALUES('Mahatma Ghandi', 'life', 'Learn as if you will live forever, live like you will die tomorrow.');
