CREATE TABLE notes 
(
    id serial primary key,
    title varchar(50) not null, 
    content text not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

CREATE TABLE attachments 
(
    id serial primary key,
    note_id integer not null,
    file_name varchar(255) not null,
    file_path varchar(255) not null,
    file_type varchar(50) not null,
    file_size integer not null,
    foreign key (note_id) references notes (id) on delete cascade
);