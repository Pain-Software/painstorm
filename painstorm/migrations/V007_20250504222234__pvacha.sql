COPY city(name, findname, country, longitude, latitude)
FROM '/data/migrations/cities.csv'
DELIMITER ','
CSV HEADER;

CREATE EXTENSION pg_cron;

SELECT cron.schedule(
  'delete_old_measurements',
  '*/30 * * * *', 
  $$DELETE FROM measurement WHERE created_at < NOW() - INTERVAL '2 days'$$
);