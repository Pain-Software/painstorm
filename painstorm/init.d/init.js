db = new Mongo().getDB("painstorm");

if (!db.getCollectionNames().includes("measurements")) {
  db.createCollection('measurements');
}

db.measurements.createIndex(
  { timestamp: 1 },
  { expireAfterSeconds: 60*60*24*5, name: "timestamp_ttl" }
);
