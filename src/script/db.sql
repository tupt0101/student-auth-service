create table users (
	user_id serial PRIMARY KEY,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	name VARCHAR ( 50 ) NOT NULL,
	created_on TIMESTAMP NOT NULL
);

create table access_tokens (
    user_id serial not null,
    access_token varchar ( 50 ) not null,
    expires bigint not null
);

create table students (
    student_id varchar ( 50 ) not null,
    name varchar ( 50 ) not null,
    class varchar ( 50 ) not null,
    created_on timestamp not null
);