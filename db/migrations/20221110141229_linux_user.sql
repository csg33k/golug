-- migrate:up
create table
    linux_user (
        linux_user_id bigint generated always as Identity,
        linux_user_name text not null unique,
        linux_distro text not null,
        primary key(linux_user_id)
    );

-- migrate:down

drop table linux_user;
