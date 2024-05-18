CREATE TABLE IF NOt EXISTS article_bodies(
    id uuid primary key not null,
    article_summaries_id text unique not null,
    body text not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);

