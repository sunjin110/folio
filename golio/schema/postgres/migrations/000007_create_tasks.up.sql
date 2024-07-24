create table if not exists tasks(
    id uuid primary key not null,
    title text not null,
    status smallint not null,
    start_time timestamptz null,
    due_time timestamptz null,
    created_at timestamptz not null, 
    updated_at timestamptz not null
);
