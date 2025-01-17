-- +goose Up
-- +goose StatementBegin
CREATE TABLE events
(
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL default current_timestamp,
    finished_at TIMESTAMP NULL,
    description TEXT NOT NULL,
    owner_id BIGINT NOT NULL,
    notify_before INT NOT NULL,
    is_notified BOOLEAN NOT NULL
);
ALTER SEQUENCE events_id_seq RESTART WITH 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events
-- +goose StatementEnd
