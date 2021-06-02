-- Create schemas

-- Create tables
CREATE TABLE IF NOT EXISTS login
(
    id SERIAL PRIMARY KEY,
    login VARCHAR(300) NOT NULL UNIQUE,
    password VARCHAR(300) NOT NULL,
    user_id INTEGER
);

CREATE TABLE IF NOT EXISTS "user"
(
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(300),
    lastname VARCHAR(300),
    username VARCHAR(300),
    mobile_phone VARCHAR(300),
    email VARCHAR(300),
    account_id INTEGER,
    updated_at DATE,
    deleted_at DATE,
    created_at DATE
);

CREATE TABLE IF NOT EXISTS location
(
    id SERIAL PRIMARY KEY,
    city VARCHAR(300),
    address VARCHAR(300),
    point Point,
    latitude VARCHAR(100),
    longitude VARCHAR(100),
    zip VARCHAR(300),
    country VARCHAR(300)
);

CREATE TABLE IF NOT EXISTS account
(
    id SERIAL PRIMARY KEY,
    type VARCHAR(300)
);

CREATE TABLE IF NOT EXISTS offer
(
    id SERIAL PRIMARY KEY,
    author_id INTEGER,
    title VARCHAR(300),
    description VARCHAR(600),
    price INTEGER,
    negociation BOOLEAN,
    category INTEGER,
    design INTEGER,
    updated_at DATE,
    deleted_at DATE,
    created_at DATE,
    location_id INTEGER,
    wheel_size INTEGER,
    size VARCHAR(2),
    condition INTEGER,
    enabled BOOLEAN,
    quantity INTEGER
);

CREATE TABLE IF NOT EXISTS offer_gear
(
    id SERIAL PRIMARY KEY,
    offer_id INTEGER,
    gear_id INTEGER
);

CREATE TABLE IF NOT EXISTS gear
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(300),
    description VARCHAR(500),
    category INTEGER
);

CREATE TABLE IF NOT EXISTS media
(
    id SERIAL PRIMARY KEY,
    thumbnail VARCHAR,
    raw VARCHAR
);

CREATE TABLE IF NOT EXISTS offer_media
(
    id SERIAL PRIMARY KEY,
    offer_id INTEGER,
    media_id INTEGER
);

CREATE TABLE IF NOT EXISTS category
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(300),
    description VARCHAR(600),
    parent_id INTEGER
);

CREATE TABLE IF NOT EXISTS offer_design
(
    id SERIAL PRIMARY KEY,
    theme VARCHAR(1000),
    color VARCHAR(300)
);

CREATE TABLE IF NOT EXISTS bike_offer
(
    id SERIAL PRIMARY KEY,
    size VARCHAR(2),
    wheel_size VARCHAR(10),
    offer_id INTEGER,
    electric_assist BOOLEAN
);

-- Create FKs
ALTER TABLE "user"
    ADD    FOREIGN KEY (account_id)
    REFERENCES account(id)
    MATCH SIMPLE
;
    
ALTER TABLE login
    ADD    FOREIGN KEY (user_id)
    REFERENCES "user"(id)
    MATCH SIMPLE
;
    
ALTER TABLE offer
    ADD    FOREIGN KEY (author)
    REFERENCES "user"(id)
    MATCH SIMPLE
;
    
ALTER TABLE offer_gear
    ADD    FOREIGN KEY (gear_id)
    REFERENCES gear(id)
    MATCH SIMPLE
;
    
ALTER TABLE offer_gear
    ADD    FOREIGN KEY (offer_id)
    REFERENCES offer(id)
    MATCH SIMPLE
;
    
ALTER TABLE offer_media
    ADD    FOREIGN KEY (media_id)
    REFERENCES media(id)
    MATCH SIMPLE
;
    
ALTER TABLE offer_media
    ADD    FOREIGN KEY (offer_id)
    REFERENCES offer(id)
    MATCH SIMPLE
;
    
ALTER TABLE gear
    ADD    FOREIGN KEY (category)
    REFERENCES category(id)
    MATCH SIMPLE
;
    
ALTER TABLE offer
    ADD    FOREIGN KEY (design)
    REFERENCES offer_design(id)
    MATCH SIMPLE
;
    
ALTER TABLE offer
    ADD    FOREIGN KEY (location_id)
    REFERENCES location(id)
    MATCH SIMPLE
;
    
ALTER TABLE bike_offer
    ADD    FOREIGN KEY (offer_id)
    REFERENCES offer(id)
    MATCH SIMPLE
;
    
ALTER TABLE offer
    ADD    FOREIGN KEY (category)
    REFERENCES category(id)
    MATCH SIMPLE
;
    
    
ALTER TABLE category
    ADD    FOREIGN KEY (parent_id)
    REFERENCES category(id)
    MATCH SIMPLE
;
    
ALTER TABLE gear
    ADD    FOREIGN KEY (category)
    REFERENCES category(id)
    MATCH SIMPLE
;
    

-- Create Indexes

INSERT INTO account(type) 
VALUES('FREE');

INSERT INTO offer_design (theme,color) 
VALUES('premium','#1389B3');

INSERT INTO offer_design (theme,color) 
VALUES('simple','#FD9055');

INSERT INTO offer_design (theme,color) 
VALUES('playground','#FF9D88');

-- Add PostGIS extension
CREATE EXTENSION postgis;