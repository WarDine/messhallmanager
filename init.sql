-- create database wardine;
\c wardine;

create sequence mass_hall_admins_seq;

create sequence mass_hall_seq;

create sequence menu_seq;

create sequence main_admin_seq;

create sequence recipe_seq;

create sequence ingredient_seq;

create table menu (
    menu_uid integer primary key not null default nextval('menu_seq'),
    recipe_uid integer,
    time_stamp timestamp not null default current_timestamp
);

create table mass_hall (
    mass_halls_uid integer primary key not null default nextval('mass_hall_seq'),
    street varchar (255) not null,
    city varchar (255) not null,
    county varchar (255) not null,
    menu_uid integer not null,
    attendance_number integer not null
    -- constraint fk_mass_hall foreign key (menu_uid) references menu (menu_uid)
);

create table mass_halls_admins (
    mass_halls_admins_uid integer primary key not null default nextval('mass_hall_admins_seq'),
    nickname varchar(255) not null,
    mass_hall_uid integer unique
    -- constraint fk_mass_halls_admins foreign key (mass_hall_uid) references mass_hall (mass_halls_uid)
);

create table main_admins (
    main_admin_uid integer primary key not null default nextval('main_admin_seq'),
    mass_hall_admin_uid integer,
    nickname varchar(255) not null
    -- constraint fk_main_admins foreign key (mass_hall_admin_uid) references mass_halls_admins (mass_halls_admins_uid)
);

create table recipe (
    recipe_uid integer primary key not null default nextval('recipe_seq'),
    mass_hall_uid integer,
    nname varchar(255) not null,
    description varchar(1000) not null,
    calories integer,
    cooking_time integer,
    instructions varchar(255),
    portions integer
    -- constraint fk_recipe foreign key (mass_hall_uid) references mass_hall (mass_halls_uid)
);

create table ingredient (
    ingredient_uid integer primary key not null default nextval('ingredient_seq'),
    ingredient_name varchar(255),
    recipe_uid integer,
    amount integer
);

create table ingredients (
    ingredient_uid integer,
    ingredient_name varchar(255)
    -- constraint fk_ingredients foreign key (ingredient_uid) references ingredient (ingredient_uid)
);