create table if not exists task_detials(
    id uuid primary key not null,
    task_id uuid unique not null,
    detail text not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);
