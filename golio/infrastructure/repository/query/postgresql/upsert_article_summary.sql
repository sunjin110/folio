insert into
    article_summaries(id, title, created_at, updated_at, tag_ids)
values
    (:id, :title, :created_at, :updated_at, :tag_ids) on conflict (id) do
update
set
    title = excluded.title,
    updated_at = excluded.updated_at,
    tag_ids = excluded.tag_ids;
