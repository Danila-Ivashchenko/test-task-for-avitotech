CREATE TABLE IF NOT EXISTS users (
    id BIGINT NOT NULL PRIMARY KEY UNIQUE,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS segments (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(200) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_segment_name ON segments(name);

CREATE TABLE IF NOT EXISTS user_in_segment (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    segment_id BIGINT NOT NULL REFERENCES segments(id)
);

CREATE UNIQUE INDEX unique_user_segment ON user_in_segment (user_id, segment_id);

CREATE TABLE IF NOT EXISTS history (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    segment_id BIGINT NOT NULL REFERENCES segments(id),
    action VARCHAR(3) CONSTRAINT history_action_check CHECK(action = 'ADD' OR action = 'DEL'),
    action_time TIMESTAMP DEFAULT now()
);

