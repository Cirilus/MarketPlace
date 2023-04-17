CREATE TABLE IF NOT EXISTS auth (
    auth_id serial PRIMARY KEY,
    username varchar(50),
    password varchar(50)
);

CREATE TABLE IF NOT EXISTS product(
    product_id serial PRIMARY KEY,
    title varchar(50),
    cost int,
    description text,
    author_id int,
    category text,
    rate float4,

    FOREIGN KEY(author_id) REFERENCES auth(auth_id)
);