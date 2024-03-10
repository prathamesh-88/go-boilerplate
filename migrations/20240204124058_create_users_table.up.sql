CREATE TYPE role_enum AS ENUM ('owner', 'admin', 'readonly');

CREATE TABLE IF NOT EXISTS users(
   username VARCHAR (50) PRIMARY KEY,
   password VARCHAR (255) NOT NULL,
   email VARCHAR (255) UNIQUE NOT NULL,
   role role_enum NOT NULL,
   is_email_verified BOOLEAN NOT NULL DEFAULT FALSE,
   created_at TIMESTAMPTZ NOT NULL,
   updated_at TIMESTAMPTZ NOT NULL
);