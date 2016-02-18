drop table if exists test cascade;
drop table if exists install cascade;
--drop table if exists machine;
drop table if exists arch;

create table arch (
	id serial unique,
	created timestamp without time zone default now(),
	name text not null
);

--create table machine (
--	id serial unique,
--	created timestamp without time zone default now(),
--	name text not null
--);

create table install (
	id serial unique, 
	created timestamp without time zone default now(),
	name text not null,
	archid int references arch (id)
); 

create table test (
	id serial unique,
	created timestamp without time zone default now(),
	snap timestamp without time zone default now(),
	tester text not null,
	name text not null,
	note text default '',
	x11 bool default false,
	status bool default false,
	upgrade bool default false,
	installid int references install (id)
	-- archid int references arch (id)
	-- machid int references machine (id)
);

CREATE OR REPLACE FUNCTION archByID(text) RETURNS integer
    AS 'select id from arch where name = $1;'
    LANGUAGE SQL
    IMMUTABLE
    RETURNS NULL ON NULL INPUT;

CREATE OR REPLACE FUNCTION getInstID(text, text) RETURNS integer
    AS 'select id from install where name = $1 and archid = archByID($2);'
    LANGUAGE SQL
    IMMUTABLE
    RETURNS NULL ON NULL INPUT;

insert into arch (name) values ('alpha');
insert into arch (name) values ('amd64');
insert into arch (name) values ('armish');
insert into arch (name) values ('armv7');
insert into arch (name) values ('aviion');
insert into arch (name) values ('hppa');
insert into arch (name) values ('i386');
insert into arch (name) values ('landisk');
insert into arch (name) values ('loongson');
insert into arch (name) values ('luna88k');
insert into arch (name) values ('macppc');
insert into arch (name) values ('octeon');
insert into arch (name) values ('sgi');
insert into arch (name) values ('socppc');
insert into arch (name) values ('sparc');
insert into arch (name) values ('sparc64');
insert into arch (name) values ('vax');
insert into arch (name) values ('zaurus');

--insert into machine (name) values ('Lenovo X240');

insert into install (name, archid) values ('install59.iso', archByID('amd64'));
insert into install (name, archid) values ('cd59.iso', archByID('amd64'));
insert into install (name, archid) values ('bsd.rd', archByID('amd64'));
insert into install (name, archid) values ('floppy59.fs', archByID('amd64'));
insert into install (name, archid) values ('install59.fs', archByID('amd64'));
insert into install (name, archid) values ('miniroot59.fs', archByID('amd64'));
insert into install (name, archid) values ('pxeboot', archByID('amd64'));

insert into install (name, archid) values ('install59.iso', archByID('i386'));
insert into install (name, archid) values ('cd59.iso', archByID('i386'));
insert into install (name, archid) values ('bsd.rd', archByID('i386'));
insert into install (name, archid) values ('floppy59.fs', archByID('i386'));
insert into install (name, archid) values ('install59.fs', archByID('i386'));
insert into install (name, archid) values ('miniroot59.fs', archByID('i386'));
insert into install (name, archid) values ('pxeboot', archByID('i386'));

insert into install (name, archid) values ('install59.iso', archByID('macppc'));

--insert into test (tester, snap, name, x11, status, installid ) values ('abieber', '2016-02-08', 'Lenovo X240', true, true, getInstID('install59.iso', 'amd64'));
--insert into test (tester, snap, name, status, installid ) values ('abieber', '2016-02-08', 'Lenovo X240', false, getInstID('floppy59.fs', 'amd64'));
--insert into test (tester, snap, name, status, installid ) values ('abieber', '2016-02-08', 'Lenovo X240', true, getInstID('pxeboot', 'amd64'));

--insert into test (tester, snap, name, x11, upgrade, status, installid ) values ('abieber', '2016-02-08', 'CR48', true, true, true, getInstID('miniroot59.fs', 'i386'));
--insert into test (tester, snap, name, x11, upgrade, status, installid ) values ('abieber', '2016-02-08', 'CR48', true, false, true, getInstID('bsd.rd', 'i386'));
