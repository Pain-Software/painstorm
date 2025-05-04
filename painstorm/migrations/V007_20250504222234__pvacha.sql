COPY city(name, findname, country, longitude, latitude)
FROM '/data/migrations/cities.csv'
DELIMITER ','
CSV HEADER;
