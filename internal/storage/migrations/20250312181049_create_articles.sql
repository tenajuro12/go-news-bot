-- +goose Up
-- +goose StatementBegin
CREATE TABLE sources
(
    id         SERIAL PRIMARY KEY,
    source_id int not null ,
    title       varchar(255) NOT NULL,
    link        varchar(255) NOT NULL,
    summary TEXT not null ,п
    published_at timestamp    NOT NULL DEFAULT NOW(),
    created_at timestamp    NOT NULL DEFAULT NOW(),
    posted_at timestamp
);

-- +goose Down
-- +goose StatementBegin
DROP TABLE sources-- +goose StatementEnd
