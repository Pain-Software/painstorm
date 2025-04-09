CREATE TABLE test_vers_lease (
    id SERIAL PRIMARY KEY,
    lessee_name VARCHAR(255) NOT NULL,
    asset_id INT NOT NULL
);
