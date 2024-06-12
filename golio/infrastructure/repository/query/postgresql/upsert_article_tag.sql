INSERT INTO
  article_tags(id, name, created_at, updated_at)
VALUES
  (:id, :name, :created_at, :updated_at) ON CONFLICT (id) DO
UPDATE
SET
  name = excluded.name,
  updated_at = excluded.updated_at;
