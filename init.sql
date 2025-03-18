create table if not exists "user"
(
    id       uuid default gen_random_uuid() not null
    constraint user_pk primary key,
    username text                           not null,
    mail     text                           not null
    );

create table if not exists leaderboard(
                                          id uuid default gen_random_uuid() not null constraint leaderboard_pk primary key ,
    name text,
    game_id uuid,
    start_date timestamptz,
    end_date timestamptz
    );

create table if not exists game(
                                   id uuid default gen_random_uuid() not null constraint leaderboard_pk primary key ,
    name text
    );

create table if not exists "group"
(
    id     uuid default gen_random_uuid() not null
    constraint group_pk primary key,
    name   text                           not null,
    icon   text
    );

create table if not exists group_leaderboard(
                                                leaderboard_id uuid,
                                                group_id uuid,
                                                primary key (leaderboard_id,group_id)
    );


alter table leaderboard
    add constraint leaderboard_game_id foreign key (game_id) references game;
alter table group_leaderboard
    add constraint leaderboard_group_id foreign key (leaderboard_id) references leaderboard;
alter table group_leaderboard
    add constraint group_leaderboard_id foreign key (group_id) references "group";




