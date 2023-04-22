CREATE TABLE IF NOT EXISTS users (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  role VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  phone_number VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);

INSERT INTO users VALUES (
  GEN_RANDOM_UUID(),
  'Joao Polidoro',
  'admin',
  'jv.polidoro@outlook.com',
  '5515997311989', 
  '$2a$10$sm4jd2vFYhG7iBVgrIq50.l.8QlSEXLNcp65WrtkJEAZDbuWjkbU2',
  NOW(),
  NOW()
);