insert into 
    tasks(id, title, status, start_time, due_time, created_at, updated_at)
values
    (:id, :title, :status, :start_time, :due_time, :created_at, :updated_at) on conflict(id) do
update
set
    title = excluded.title,
    status = excluded.status,
    start_time = excluded.start_time,
    due_time = excluded.due_time,
    updated_at = excluded.updated_at
