CREATE TABLE IF NOT EXISTS quizes(
  id serial PRIMARY KEY,
  author_id serial,
  quiz_hash character varying(64),
  title character varying(128),
  questions jsonb,
  active bool default false,
  created_at timestamp without time zone default now(),
  updated_at timestamp without time zone default now()
);