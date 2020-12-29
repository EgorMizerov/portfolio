CREATE TABLE works
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255) not null,
    date_up varchar(255) not null,
    tag varchar(255) not null,
    img varchar(500) not null,
    url varchar(255) not null
);