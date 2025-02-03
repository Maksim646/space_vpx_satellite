BEGIN;


UPDATE schema_migrations SET dirty=false;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS main_components
(
    id                       BIGSERIAL PRIMARY KEY,
    name                     VARCHAR(255),
    updated_at               TIMESTAMP WITHOUT TIME ZONE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

DROP TABLE IF EXISTS main_components;

CREATE TABLE IF NOT EXISTS users
(
    id                       uuid DEFAULT uuid_generate_v4(),
    name                     VARCHAR(255),
    email                    VARCHAR(255) UNIQUE,
    password_hash            VARCHAR(255),
    updated_at               TIMESTAMP WITHOUT TIME ZONE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS projects
(
    id                       uuid DEFAULT uuid_generate_v4(),
    name                     VARCHAR(255),
    user_id                  VARCHAR(255),
    chassis_id               BIGINT,
    updated_at               TIMESTAMP WITHOUT TIME ZONE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS chassis
(
    id                                  BIGSERIAL PRIMARY KEY,
    name                                VARCHAR(255) UNIQUE,
    slots                               BIGINT,
    size                                VARCHAR(255),
    max_operating_temperature           NUMERIC,
    min_operating_temperature           NUMERIC,
    max_non_operating_temperature       NUMERIC,
    min_non_operating_temperature       NUMERIC,
    overload                            NUMERIC,
    max_vibration_sine                  NUMERIC,
    min_vibration_sine                  NUMERIC,
    max_vibration_random                NUMERIC,
    min_vibration_random                NUMERIC,
    axes                                BIGINT,
    shock_response_spectrum             NUMERIC,
    peak_overload_spectrum_1            NUMERIC,
    peak_overload_spectrum_2            NUMERIC,
    peak_frequency_spectrum_1           NUMERIC,
    peak_frequency_spectrum_2           NUMERIC,
    length                              NUMERIC,
    width                               NUMERIC,
    height                              NUMERIC,
    weight                              NUMERIC,
    power_handling_capability_per_board NUMERIC,
    temperature_per_board               NUMERIC,
    updated_at                          TIMESTAMP WITHOUT TIME ZONE,
    created_at                          TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS cube_sat_frame
(
    id                                  BIGSERIAL PRIMARY KEY,
    length                              NUMERIC,
    width                               NUMERIC,
    height                              NUMERIC,
    weight                              BIGINT,
    operating_temperature_min           BIGINT,
    operating_temperature_max           BIGINT,
    mechanical_vibration                BIGINT,
    mechanical_shock                    BIGINT,
    link                                VARCHAR(255),
    updated_at                          TIMESTAMP WITHOUT TIME ZONE,
    created_at                          TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS admins
(
    id                       uuid DEFAULT uuid_generate_v4(),
    name                     VARCHAR(255),
    email                    VARCHAR(255) UNIQUE,
    password_hash            VARCHAR(255),
    updated_at               TIMESTAMP WITHOUT TIME ZONE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);


COMMIT;


