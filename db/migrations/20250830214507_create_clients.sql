-- +goose Up
-- +goose StatementBegin
create table if not exists clients
(
    id         bigserial primary key,
    fio        text        not null,
    phone      text unique not null,
    email      text unique not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists clients;
-- +goose StatementEnd
