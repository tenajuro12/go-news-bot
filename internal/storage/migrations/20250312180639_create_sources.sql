-- +goose Up
-- +goose StatementBegin
CREATE TABLE sources
(
    id         SERIAL PRIMARY KEY,
    name       varchar(255) NOT NULL,
    feed_url        varchar(255) NOT NULL,
    created_at timestamp    NOT NULL DEFAULT NOW(),
    updated_at timestamp    NOT NULL DEFAULT NOW()
);

-- +goose Down
-- +goose StatementBegin
DROP TABLE sources-- +goose StatementEnd
