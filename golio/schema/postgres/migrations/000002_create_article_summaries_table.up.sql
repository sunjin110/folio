CREATE TABLE IF NOT EXISTS article_summaries(
    id uuid primary key not null,
    title text not null,
    created_at timestamp not null, 
    updated_at timestamp not null
);
