insert into 
    task_details(id, task_id, detail, created_at, updated_at)
values
    (:id, :taks_id, :detail, :created_at, :updated_at) on conflict(id) do
update
set
    task_id = excluded.task_id,
    detail = excluded.detail,
    created_at = excluded.created_at,
    updated_at = excluded.updated_at
