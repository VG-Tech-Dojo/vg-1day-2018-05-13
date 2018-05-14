-- +migrate Up
CREATE TABLE message (
    id INTEGER NOT NULL PRIMARY KEY,
    body TEXT NOT NULL DEFAULT "",
    username TEXT NOT NULL DEFAULT "",
    msgtype INTEGER NOT NULL DEFAULT 0,
    created TIMESTAMP NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    updated TIMESTAMP NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);

-- +migrate Down
DROP TABLE message;
