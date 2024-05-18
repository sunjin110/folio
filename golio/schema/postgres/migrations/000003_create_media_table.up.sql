create table if not exists media(
    id uuid primary key not null,
    path text not null,
    file_type text not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);
