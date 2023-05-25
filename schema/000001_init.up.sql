CREATE TABLE users
(
    id serial primary key,
    username varchar(255) not null,
    password_hash varchar(255) not null,
    created_at timestamp
);

CREATE TABLE users_data
(
    id serial primary key,
    user_id int references users(id) on delete cascade not null,
    age int not null,
    sex boolean not null,
    weight int not null,
    height int not null,
    goal varchar(255) not null,
    place varchar(255) not null,
    calories_intake int not null
);

CREATE TABLE sessions
(
    id serial primary key,
    user_id int references users(id) on delete cascade not null,
    session varchar(255) not null,
    created_at timestamp
);

CREATE TABLE workouts
(
    id serial primary key,
    workout varchar(255) not null,
    calories int not null
);

CREATE TABLE workouts_archive
(
    id serial primary key,
    user_id int references users(id),
    workout_id int references workouts(id)
);
