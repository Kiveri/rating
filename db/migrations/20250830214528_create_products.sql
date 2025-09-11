-- +goose Up
-- +goose StatementBegin
create table if not exists products
(
    id bigserial primary key,
    name text not null,
    article text unique not null,
    price bigint not null,
    product_type smallint not null,
    average_rate double precision not null,
    feedbacks_count bigint not null default 0,
    created_at timestamptz not null,
    updated_at timestamptz not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists products;
-- +goose StatementEnd
