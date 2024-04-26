CREATE TABLE IF NOT EXISTS article_summaries(
    id text primary key not null,
    title text not null,
    created_at integer not null, 
    updated_at integer not null
);
