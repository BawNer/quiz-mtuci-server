CREATE TABLE IF NOT EXISTS reviewers(
    id serial primary key,
    quiz_id serial,
    answers jsonb,
    closed_at timestamp without time zone default now()
);