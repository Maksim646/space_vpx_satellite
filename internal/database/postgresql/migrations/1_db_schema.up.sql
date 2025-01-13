BEGIN;


-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS main_components
(
    id                       BIGSERIAL PRIMARY KEY,
    name                     VARCHAR(255),
    updated_at               TIMESTAMP WITHOUT TIME ZONE

);
COMMIT;


