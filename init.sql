create table if not exists "user"
(
    id       uuid default gen_random_uuid() not null
    constraint user_pk primary key,
    username text                           not null,
    mail     text                           not null,
    group_id uuid
    );

create table if not exists leaderboard(
    id uuid default gen_random_uuid() not null constraint leaderboard_pk primary key ,
    name text,
    start_date timestamptz,
    end_date timestamptz
    );

create table if not exists game(
    id uuid default gen_random_uuid() not null constraint leaderboard_pk primary key ,
    name text
    leaderboard_id uuid
    );

create table if not exists "group"
(
    id     uuid default gen_random_uuid() not null constraint group_pk primary key,
    name   text                           not null,
    icon   text
    );

alter table game add constraint game_leaderboard_id foreign key (leaderboard_id) references leaderboard;
alter table "user" add constraint user_group_id foreign key (group_id) references "group";




