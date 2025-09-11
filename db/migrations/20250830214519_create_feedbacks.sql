-- +goose Up
-- +goose StatementBegin
create table if not exists feedbacks
(
    id         bigserial primary key,
    rate       smallint    not null,
    comment    text,
    product_id bigint      not null,
    client_id  bigint      not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists feedbacks;
-- +goose StatementEnd
