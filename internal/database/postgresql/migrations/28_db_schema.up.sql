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

CREATE TABLE IF NOT EXISTS cube_sat_projects
(
    id                            uuid DEFAULT uuid_generate_v4(),
    name                          VARCHAR(255),
    user_id                       VARCHAR(255),
    size                          BIGINT,
    frame_name                    VARCHAR(255),
    solar_panel_side_name         VARCHAR(255),
    solar_panel_top_name          VARCHAR(255),
    board_computing_module_name VARCHAR(255),
    power_system_name             VARCHAR(255),
    vhf_antenna_system_name       VARCHAR(255),
    vhf_transceiver_name          VARCHAR(255),
    updated_at                    TIMESTAMP WITHOUT TIME ZONE,
    created_at                    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
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
    name                                VARCHAR(255) UNIQUE,
    length                              NUMERIC,
    width                               NUMERIC,
    height                              NUMERIC,
    weight                              NUMERIC,
    size                                BIGINT,
    interface                           VARCHAR(255),
    operating_temperature_min           BIGINT,
    operating_temperature_max           BIGINT,
    mechanical_vibration                BIGINT,
    mechanical_shock                    BIGINT,
    link                                VARCHAR(255),
    updated_at                          TIMESTAMP WITHOUT TIME ZONE,
    created_at                          TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS cube_sat_solar_panel_side
(
    id                                  BIGSERIAL PRIMARY KEY,
    name                                VARCHAR(255) UNIQUE,
    length                              NUMERIC,
    width                               NUMERIC,
    height                              NUMERIC,
    weight                              NUMERIC,
    interface                           VARCHAR(255),
    voc                                 NUMERIC,
    isc                                 NUMERIC,
    vmp                                 NUMERIC,
    imp                                 NUMERIC,
    efficiency                          BIGINT,
    coil_area                           NUMERIC,
    coil_resistance                     BIGINT,
    max_operating_temperature           NUMERIC,
    min_operating_temperature           NUMERIC,
    mechanical_vibration                BIGINT,
    mechanical_shock                    BIGINT,
    updated_at                          TIMESTAMP WITHOUT TIME ZONE,
    created_at                          TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')

);

CREATE TABLE IF NOT EXISTS cube_sat_solar_panel_top (
    id                                  BIGSERIAL PRIMARY KEY,
    name                                VARCHAR(255) UNIQUE,       
    length                              NUMERIC,                   
    width                               NUMERIC,                    
    height                              NUMERIC,                    
    weight                              NUMERIC,                                 
    interface                           VARCHAR(255),               
    voc                                 NUMERIC,                    
    isc                                 NUMERIC,                    
    vmp                                 NUMERIC,                   
    imp                                 NUMERIC,                    
    efficiency                          NUMERIC,                    
    coil_area                           NUMERIC,
    coil_resistance                     BIGINT,                  
    min_operating_temperature           NUMERIC,                    
    max_operating_temperature           NUMERIC,                                       
    mechanical_vibration                NUMERIC,                   
    mechanical_shock                    NUMERIC,                   
    updated_at                          TIMESTAMP WITHOUT TIME ZONE,
    created_at                          TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS cube_sat_power_system
(
    id                                  BIGSERIAL PRIMARY KEY,
    name                                VARCHAR(255) UNIQUE,
    length                              NUMERIC,
    width                               NUMERIC,
    height                              NUMERIC,
    weight                              NUMERIC,
    solar_panel_channels                BIGINT,
    solar_panels_type                   VARCHAR(255),
    solar_panel_voltage_min             NUMERIC,
    solar_panel_voltage_max             NUMERIC,
    solar_panel_current_per_channel_max NUMERIC,
    total_current_of_solar_panels_max   NUMERIC,
    output_channels                     BIGINT,
    system_bus_voltage_solar_panels     NUMERIC,
    max_system_bus_voltage_output_channels  NUMERIC,
    min_system_bus_voltage_output_channels  NUMERIC,
    current_output_channels_max         NUMERIC,
    total_output_current                NUMERIC,
    data_interface                      VARCHAR(255),
    max_operating_temperature           NUMERIC,
    min_operating_temperature           NUMERIC,
    mechanical_vibration                BIGINT,
    mechanical_shock                    BIGINT,
    updated_at                          TIMESTAMP WITHOUT TIME ZONE,
    created_at                          TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS board_computing_module
(
    id                                  BIGSERIAL PRIMARY KEY,
    name                                VARCHAR(255) UNIQUE,
    length                              NUMERIC,
    width                               NUMERIC,
    height                              NUMERIC,
    weight                              NUMERIC,
    max_supply_voltage                  NUMERIC,
    min_supply_voltage                  NUMERIC,
    power_consumption                   NUMERIC,
    interface                           VARCHAR(255),
    max_operating_temperature           NUMERIC,
    min_operating_temperature           NUMERIC,
    mechanical_vibration                BIGINT,
    mechanical_shock                    BIGINT,
    data_bus                            VARCHAR(255),
    updated_at                          TIMESTAMP WITHOUT TIME ZONE,
    created_at                          TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS vhf_transceiver
(
    id                                  BIGSERIAL PRIMARY KEY,
    name                                VARCHAR(255) UNIQUE,
    length                              NUMERIC,
    width                               NUMERIC,
    height                              NUMERIC,
    weight                              NUMERIC,
    supply_voltage                      NUMERIC,
    power_consumption                   NUMERIC,
    interface                           VARCHAR(255),
    operating_frequency                 NUMERIC,
    output_power                        NUMERIC,  -- Дб
    sensitivity_receiver                NUMERIC,
    max_operating_temperature           NUMERIC,
    min_operating_temperature           NUMERIC,
    mechanical_vibration                BIGINT,
    mechanical_shock                    BIGINT,
    updated_at                          TIMESTAMP WITHOUT TIME ZONE,
    created_at                          TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS vhf_antenna_system
(
    id                                  BIGSERIAL PRIMARY KEY,
    name                                VARCHAR(255) UNIQUE,
    length                              NUMERIC,
    width                               NUMERIC,
    height                              NUMERIC,
    weight                              NUMERIC,
    interface                           VARCHAR(255),
    frequency                           NUMERIC,
    max_operating_temperature           NUMERIC,
    min_operating_temperature           NUMERIC,
    mechanical_vibration                BIGINT,
    mechanical_shock                    BIGINT,
    updated_at                          TIMESTAMP WITHOUT TIME ZONE,
    created_at                          TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS admins
(
    id                       uuid DEFAULT uuid_generate_v4(),
    name                     VARCHAR(255),
    email                    VARCHAR(255) UNIQUE,
    role                     VARCHAR(50) NOT NULL DEFAULT 'admin',
    password_hash            VARCHAR(255),
    updated_at               TIMESTAMP WITHOUT TIME ZONE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);


COMMIT;


