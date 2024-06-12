create table if not exists article_tags(
    id uuid primary key not null,
    name text not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);
