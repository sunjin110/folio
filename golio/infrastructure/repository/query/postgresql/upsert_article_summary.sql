insert into
    article_summaries(id, title, created_at, updated_at)
values
    (:id, :title, :created_at, :updated_at) on conflict (id) do
update
set
    title = excluded.title,
    updated_at = excluded.updated_at;
    