CREATE TABLE article_bodies(
    id text primary key not null,
    article_summaries_id text unique not null,
    body text not null,
    created_at integer,
    updated_at integer
);
