CREATE TABLE IF NOT EXISTS profile (
  id   BIGSERIAL PRIMARY KEY,
  user_id BIGINT UNIQUE NOT NULL,
  name VARCHAR(255) NOT NULL,
  bio TEXT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
