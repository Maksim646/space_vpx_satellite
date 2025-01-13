BEGIN;


-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
UPDATE schema_migrations SET dirty=false;

CREATE TABLE IF NOT EXISTS main_components
(
    id                       BIGSERIAL PRIMARY KEY,
    name                     VARCHAR(255),
    updated_at               TIMESTAMP WITHOUT TIME ZONE

);

CREATE TABLE IF NOT EXISTS users
(
    id                       uuid DEFAULT uuid_generate_v4(),
    name                     VARCHAR(255),
    email                    VARCHAR(255) UNIQUE,
    password_hash            VARCHAR(255) UNIQUE,
    updated_at               TIMESTAMP WITHOUT TIME ZONE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);


COMMIT;


