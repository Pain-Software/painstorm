DROP TABLE city CASCADE;
DROP TABLE weather CASCADE;
DROP TABLE measurement CASCADE;
DROP TABLE weather_in_measurement CASCADE;
DROP TABLE stats CASCADE;

CREATE TABLE city (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    findName VARCHAR(255),
    country VARCHAR(100),
    longitude DECIMAL(9,6),
    latitude DECIMAL(9,6)
);

CREATE TABLE weather (
    id SERIAL PRIMARY KEY,
    main VARCHAR(100),
    description VARCHAR(100),
    icon VARCHAR(100)
);

CREATE TABLE measurement (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    id_city BIGINT NOT NULL,
    min_temperature DOUBLE PRECISION,
    max_temperature DOUBLE PRECISION,
    temperature DOUBLE PRECISION,
    humidity INTEGER CHECK (humidity BETWEEN 0 AND 100),
    pressure DOUBLE PRECISION,
    sea_level DOUBLE PRECISION,
    ground_level DOUBLE PRECISION,
    wind_speed DOUBLE PRECISION,
    wind_degrees DOUBLE PRECISION,
    rain_intensity DOUBLE PRECISION,
    CONSTRAINT fk_city FOREIGN KEY (id_city) REFERENCES city (id) ON DELETE CASCADE
);

CREATE TABLE weather_in_measurement (
    id_weather INTEGER NOT NULL,
    id_measurement INTEGER NOT NULL,
    PRIMARY KEY (id_weather, id_measurement),
    CONSTRAINT fk_wim_weather FOREIGN KEY (id_weather) REFERENCES weather (id) ON DELETE CASCADE,
    CONSTRAINT fk_wim_measurement FOREIGN KEY (id_measurement) REFERENCES measurement (id) ON DELETE CASCADE
);



CREATE INDEX idx_measurement_timestamp ON measurement(timestamp);
CREATE INDEX idx_measurement_rain ON measurement(rain_intensity);
CREATE INDEX idx_measurement_temperature ON measurement(temperature);
CREATE INDEX idx_measurement_weather ON weather(main);
-- zamyslet se pak ještě nad tímhle - neznám myšlenku zatím findName a name
-- CREATE INDEX idx_measurement_city_timestamp ON measurement(city_name, timestamp);
