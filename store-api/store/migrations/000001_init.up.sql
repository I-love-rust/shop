-- Create the table 'users'
CREATE TABLE IF NOT EXISTS users (
    id          SERIAL          PRIMARY KEY,
    username    VARCHAR(50)     NOT NULL UNIQUE
    CONSTRAINT CH_user_name CHECK (LENGTH(username) >= 3),
    name        VARCHAR(100)    NOT NULL,
    avatar      VARCHAR(200)    DEFAULT NULL,
    bio         VARCHAR(200)    DEFAULT NULL,
    email       VARCHAR(100)    NOT NULL UNIQUE
    CONSTRAINT CH_user_email CHECK (email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    password    VARCHAR(100)    NOT NULL,
    role        SMALLINT        NOT NULL DEFAULT 1,
    created_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);

-- Create insert user procedure
CREATE OR REPLACE PROCEDURE create_user(
    p_username  VARCHAR(50),
    p_email     VARCHAR(100),
    p_password  VARCHAR(50)
) LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO Users (username, name, email, password)
    VALUES (p_username, p_username, p_email, p_password);
    EXCEPTION
        WHEN unique_violation THEN
            RAISE EXCEPTION 'The email or username already exists.';
        WHEN check_violation THEN
            RAISE EXCEPTION 'One or more input parameters violate the constraints.';
END;
$$;

-- Create get user by email or username func
create or replace function get_user_by_email_or_username(in email_or_username varchar)
returns
    table(
        p_id int,
        p_username varchar,
        p_name varchar,
        p_avatar varchar,
        p_bio varchar,
        p_email varchar,
        p_password varchar,
        p_role smallint,
        p_created_at timestamp
    )
language plpgsql
as $$
BEGIN
    RETURN QUERY
    SELECT id, username, name, avatar, bio, email, password, role, created_at
    FROM users
    WHERE email = email_or_username OR username = email_or_username;
END;
$$
;

-- Create get user by id func
create or replace function get_user_by_id(in search_id int)
returns
    table(
        p_id int,
        p_username varchar,
        p_name varchar,
        p_avatar varchar,
        p_bio varchar,
        p_email varchar,
        p_password varchar,
        p_role smallint,
        p_created_at timestamp
    )
language plpgsql
as $$
BEGIN
    RETURN QUERY
    SELECT id, username, name, avatar, bio, email, password, role, created_at
    FROM users
    WHERE id = search_id;
END;
$$
;

-- Create get user by username func
create or replace function get_user_by_username(in v_username varchar)
returns
    table(
        p_id int,
        p_username varchar,
        p_name varchar,
        p_avatar varchar,
        p_bio varchar,
        p_email varchar,
        p_password varchar,
        p_role smallint,
        p_created_at timestamp
    )
language plpgsql
as $$
BEGIN
    RETURN QUERY
    SELECT id, username, name, avatar, bio, email, password, role, created_at
    FROM users
    WHERE username = v_username;
END;
$$
;


-- Create the table 'sessions'
CREATE TABLE IF NOT EXISTS sessions (
	id				UUID		PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         INT         NOT NULL REFERENCES users(id),
    refresh_token   VARCHAR(40) NOT NULL,
    created_at      TIMESTAMP   NOT NULL DEFAULT NOW(),
    expires_at      TIMESTAMP   NOT NULL DEFAULT (current_date + interval '7 days')
);

-- Create insert session procedure
CREATE OR REPLACE PROCEDURE create_session(
    p_user_id       INT,
    p_refresh_token VARCHAR(40)
) LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO sessions (user_id, refresh_token)
    VALUES (p_user_id, p_refresh_token);
END;
$$;

-- Update refresh token procedure
create or replace function
    update_session(p_new_token varchar(40), p_old_token varchar(40))
returns int
as $$
DECLARE
    p_user_id INT;
BEGIN
    IF EXISTS (SELECT 1 FROM sessions WHERE refresh_token = p_old_token) THEN
        UPDATE sessions
        SET refresh_token = p_new_token,
            expires_at = (current_date + interval '7 days')
        WHERE refresh_token = p_old_token
        RETURNING user_id INTO p_user_id;
        RETURN p_user_id;
    ELSE
        RAISE EXCEPTION 'Session not found with token: %', p_old_token;
    END IF;
END;
$$
language plpgsql
;

-- Create procedure to delete expired sessions
CREATE OR REPLACE PROCEDURE delete_expired_sessions()
LANGUAGE plpgsql AS $$
BEGIN
    DELETE FROM sessions
    WHERE expires_at < NOW();
END;
$$;


-- Create the table 'articles'
CREATE TABLE IF NOT EXISTS products (
    id          SERIAL          PRIMARY KEY,
    name            VARCHAR(100)    NOT NULL,
    path            VARCHAR(105)    NOT NULL,
    description     TEXT            NOT NULL,
    tags            TEXT[],
    price           NUMERIC         NOT NULL,
    preview         VARCHAR(200),
    characteristics JSONB           NOT NULL
);

-- Create insert article procedure
CREATE OR REPLACE PROCEDURE create_product(
    p_name          VARCHAR(100),
    p_path          VARCHAR(105),
    p_description   TEXT,
    p_tags          TEXT[],
    p_price       	NUMERIC,
    p_preview       VARCHAR(100),
    p_charac        JSONB
) LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO products (name, path, description, tags, price, preview, characteristics)
    VALUES (p_name, p_path, p_description, p_tags, p_price, p_preview, p_charac);
END;
$$;

create or replace function get_products_page(page_number int)
returns
    table(
        p_id          INT,
        p_name        VARCHAR(100),
        p_path        VARCHAR(105),
        p_description TEXT,
        p_tags        TEXT[],
        p_price       NUMERIC,
        p_preview     VARCHAR(200),
        p_charac      JSONB
    )
as $$
DECLARE
    offset_value INT;
BEGIN
    IF page_number < 1 THEN
        page_number := 1;
    END IF;
    
    offset_value := (page_number - 1) * 20;

    RETURN QUERY
    SELECT
        id AS p_id,
        name AS p_name,
        path AS p_path,
        description AS p_description,
        tags AS p_tags,
        price AS p_price,
        preview AS p_preview,
        characteristics AS p_charac
    FROM
        products
    ORDER BY
        id
    LIMIT 20
    OFFSET offset_value;
END;
$$
language plpgsql
;

create or replace function
    search_product(p_last int, p_tags text[] = null, p_text varchar = null)
returns setof products
language plpgsql
as $$
BEGIN
    RETURN QUERY
    SELECT *
    FROM products
    WHERE (tags @> p_tags OR p_tags IS NULL)
      AND (name ILIKE '%' || p_Text || '%' OR p_Text IS NULL)
      AND (description ILIKE '%' || p_Text || '%' OR p_Text IS NULL)
      AND id <= (SELECT (abs(p_last - id)) FROM products ORDER BY 1 DESC LIMIT 1);
END;
$$
;


-- Create get article by path func
create or replace function get_product_by_path(in v_path varchar)
returns
    table(
        p_id          INT,
        p_name        VARCHAR(100),
        p_path        VARCHAR(105),
        p_description TEXT,
        p_tags        TEXT[],
        p_price       NUMERIC,
        p_preview     VARCHAR(200),
        p_charac      JSONB
    )
as $$
BEGIN
    RETURN QUERY
    SELECT id, name, path, description, tags, price, preview, characteristics
    FROM products
    WHERE path = v_path;
END;
$$
language plpgsql
;

create or replace function get_product_by_id(in v_id int)
returns
    table(
        p_id          INT,
        p_name        VARCHAR(100),
        p_path        VARCHAR(105),
        p_description TEXT,
        p_tags        TEXT[],
        p_price       NUMERIC,
        p_preview     VARCHAR(200),
        p_charac      JSONB
    )
as $$
BEGIN
    RETURN QUERY
    SELECT id, name, path, description, tags, price, preview, characteristics
    FROM products
    WHERE id = v_id;
END;
$$
language plpgsql
;


CREATE TABLE IF NOT EXISTS purchase (
    id          SERIAL      PRIMARY KEY,
    user_id     INT     	NOT NULL REFERENCES users(id),
    product_id  INT     	NOT NULL REFERENCES products(id),
    created_at  TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE PROCEDURE register_purchase (
    p_user_id       INT,
    p_product_id    INT
) LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO purchase (user_id, product_id)
    VALUES (p_user_id, p_product_id);
END;
$$;
