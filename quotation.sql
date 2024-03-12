DROP TABLE IF EXISTS quotations;

CREATE TABLE quotations(
    quotation_id serial PRIMARY KEY,
    insertion_date date NOT NULL DEFAULT NOW(),
    author_name text NOT NULL,
    category text NOT NULL,
    quote text NOT NULL
);


-- insert a quote for testing purposes
INSERT INTO quotations(author_name, category, quote)
VALUES('Mahatma Ghandi', 'life', 'Learn as if you will live forever, live like you will die tomorrow.');


