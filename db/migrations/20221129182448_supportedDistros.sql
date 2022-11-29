-- migrate:up
CREATE TABLE public.linux_distro (
    linux_distro text unique not null
);

insert into public.linux_distro (linux_distro) values('Fedora');
insert into public.linux_distro (linux_distro) values('SUSE');
insert into public.linux_distro (linux_distro) values('Ubuntu');
insert into public.linux_distro (linux_distro) values('Gentoo');

alter table public.linux_user add constraint fk_distro_name FOREIGN KEY (linux_distro) REFERENCES public.linux_distro(linux_distro);

-- migrate:down
alter table public.linux_user drop constraint fk_distro_name;
drop table public.linux_distro;
