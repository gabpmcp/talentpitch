CREATE TABLE IF NOT EXISTS events (
    user_id UUID,
    id SERIAL,
    event_name VARCHAR(255) NOT NULL,
    event_data JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (user_id, id)
);