-- +migrate Up
ALTER TABLE message ADD COLUMN parent_message_id INTEGER;

-- +migrate Down
