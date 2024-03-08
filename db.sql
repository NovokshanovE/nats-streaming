CREATE DATABASE nats_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';

REATE TABLE public.message (
    id character varying NOT NULL,
    "order" jsonb
);

CREATE ROLE tm_admin LOGIN ENCRYPTED PASSWORD 'admin';
GRANT  select, insert, update, delete on message to tm_admin;
