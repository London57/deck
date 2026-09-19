create schema social;

create schema social;

create table social.profiles(
	id uuid primary key default gen_random_uuid(),
    version bigint not null default 1 /* оптимистичная блокировка */,
	email varchar(320) not null unique /* TODO: создать индекс */,
    phone_number varchar(15) check (
        phone_number ~ '^+[0-9]+$'
        and
        char_length(phone_number) between 10 and 15
    ),
	username varchar(30) not null unique check (char_length(username) between 6 and 30),
	password varchar(250) not null,
	name varchar(50) not null,
	birtday date not null,
	gender smallint not null,
	longitude real,
	latitude real,
	like_ttl smallint not null default 7 check (like_ttl between 1 and 30)
);

create type reaction_type as enum ("like", "dislike")

create table social.reactions(
	id uuid primary key default gen_random_uuid()
	source uuid references social.profiles(id) on delete cascade
	destination uuid references social.profiles(id) on delete cascade
	reaction reaction_type not null
	expires_at timestamp not null /* user_like_ttl+time.now() */
);