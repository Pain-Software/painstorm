-- Create table for CITY
CREATE TABLE city (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    findName VARCHAR(255),
    country VARCHAR(100),
    longtitude DOUBLE PRECISION,
    lattitude DOUBLE PRECISION
);

-- Create table for STATS
CREATE TABLE stats (
    id INTEGER PRIMARY KEY,
    min_temperature DOUBLE PRECISION,
    max_temperature DOUBLE PRECISION,
    temperature DOUBLE PRECISION,
    humidity INTEGER CHECK (humidity BETWEEN 0 AND 100),
    pressure INTEGER,
    sea_level DOUBLE PRECISION,
    ground_level DOUBLE PRECISION,
    wind_speed DOUBLE PRECISION,
    wind_degrees DOUBLE PRECISION
);

-- Create table for WEATHER
CREATE TABLE weather (
    id INTEGER PRIMARY KEY,
    main TEXT,
    description TEXT,
    icon TEXT
);

-- Create table for MEASUREMENT.
-- This table references city and stats.
-- Note: Although the provided mermaid schema shows an "id_weather" FK in MEASUREMENT,
-- a separate join table WEATHER_IN_MEASUREMENT is used to associate multiple weather records.
CREATE TABLE measurement (
    timestamp TIMESTAMP NOT NULL,
    id_city BIGINT NOT NULL,
    id_stat INTEGER NOT NULL,
    -- Optionally, if you still want to include an id_weather FK here uncomment below:
    -- id_weather INTEGER,
    CONSTRAINT fk_city
      FOREIGN KEY (id_city)
      REFERENCES city (id)
      ON DELETE CASCADE,
    CONSTRAINT fk_stats
      FOREIGN KEY (id_stat)
      REFERENCES stats (id)
      ON DELETE CASCADE
    -- If including id_weather:
    -- ,CONSTRAINT fk_weather
    --   FOREIGN KEY (id_weather)
    --   REFERENCES weather (id)
    --   ON DELETE SET NULL
);

-- Create join table for associating WEATHER and MEASUREMENT (many-to-many)
CREATE TABLE weather_in_measurement (
    id_weather INTEGER NOT NULL,
    id_measurement TIMESTAMP NOT NULL,
    -- You may choose to add additional unique measurement identifier if needed.
    PRIMARY KEY (id_weather, id_measurement),
    CONSTRAINT fk_wim_weather
      FOREIGN KEY (id_weather)
      REFERENCES weather (id)
      ON DELETE CASCADE,
    CONSTRAINT fk_wim_measurement
      FOREIGN KEY (id_measurement)
      REFERENCES measurement (timestamp)
      ON DELETE CASCADE
);
