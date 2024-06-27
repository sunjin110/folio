CREATE TABLE article_bodies (
 id UUID,
 article_summaries_id TEXT,
 body TEXT,
 created_at TIMESMAPTZ,
 updated_at TIMESMAPTZ
);

CREATE TABLE article_summaries (
 id UUID,
 title TEXT,
 created_at TIMESMAPTZ,
 updated_at TIMESMAPTZ,
 tag_ids TEXT[]
);

CREATE TABLE article_tags (
 id UUID,
 name TEXT,
 created_at TIMESMAPTZ,
 updated_at TIMESMAPTZ
);

CREATE TABLE media (
 id UUID,
 path TEXT,
 file_type TEXT,
 created_at TIMESMAPTZ,
 updated_at TIMESMAPTZ
);

CREATE TABLE schema_migrations (
 version BIGINT,
 dirty BOOLEAN
);