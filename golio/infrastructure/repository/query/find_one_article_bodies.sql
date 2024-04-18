select
    *
from
    article_bodies
where
    article_summaries_id = ?
limit
    1;
