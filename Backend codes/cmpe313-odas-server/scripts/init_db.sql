create table "users"
(
    "user_id"         bigserial primary key,
    "name"            varchar not null,
    "surname"         varchar not null,
    "email"           varchar not null unique,
    "username"        varchar not null,
    "gender"          varchar not null,
    "phone"           varchar not null,
    "address"         varchar not null,
    "date_of_birth"   timestamptz,
    "password_hash"   varchar not null,
    "title"           varchar null,
    "expertise"       varchar null,
    "hospital"        varchar null,
    "room"            varchar null,
    "privilege_level" bigint null,
    "role"            varchar not null
);

create table "appointments"
(
    "appointment_id" bigserial primary key,
    "title"          varchar not null,
    "patient_id"     bigserial references users (user_id),
    "doctor_id"      bigserial references users (user_id),
    "date"           timestamptz,
    "hospital"       varchar not null,
    "room"           varchar not null
);

create table "available_times"
(
    "available_times_id" bigserial primary key,
    "doctor_id"          bigserial references users (user_id),
    "date"               timestamptz,
    "hospital"           varchar not null,
    "room"               varchar not null
);
