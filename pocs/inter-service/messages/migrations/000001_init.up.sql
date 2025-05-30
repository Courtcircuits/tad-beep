CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS messages (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  content text NOT NULL,
  channel_id varchar(255) NOT NULL,
  owner_id varchar(255) NOT NULL,
  created_at varchar(255) NOT NULL
);
