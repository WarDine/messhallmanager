-- create database wardine;
\c wardine;

create sequence messhall_admins_seq;
create sequence messhall_seq;
create sequence menu_seq;
create sequence main_admin_seq;
create sequence recipe_seq;
create sequence ingredient_seq;

create table menu (
    menu_uid varchar (255) not null,
    recipe_uid varchar (255) not null,
    time_stamp timestamp not null default current_timestamp
);

create table messhall (
    messhalls_uid varchar (255) primary key not null default nextval('messhall_seq'),
    street varchar (255) not null,
    city varchar (255) not null,
    county varchar (255) not null,
    menu_uid varchar (255) not null,
    status varchar (255) not null,
    attendance_number integer not null
    -- constraint fk_messhall foreign key (menu_uid) references menu (menu_uid)
);

create table messhalls_admins (
    messhalls_admins_uid varchar (255) primary key not null default nextval('messhall_admins_seq'),
    nickname varchar(255) not null,
    messhall_uid varchar (255) unique 
-- constraint fk_messhalls_admins foreign key (messhall_uid) references messhall (messhalls_uid)
);

create table main_admins (
    main_admin_uid varchar (255) primary key not null default nextval('main_admin_seq'),
    messhall_admin_uid varchar (255),
    nickname varchar(255) not null
    -- constraint fk_main_admins foreign key (messhall_admin_uid) references messhalls_admins (messhalls_admins_uid)
);

create table recipe (
    recipe_uid varchar (255) primary key not null default nextval('recipe_seq'),
    messhall_uid varchar (255),
    nname varchar(255) not null,
    description varchar(1000) not null,
    calories integer,
    cooking_time integer,
    instructions varchar(255),
    portions integer
    -- constraint fk_recipe foreign key (messhall_uid) references messhall (messhalls_uid)
);

create table stock_ingredients (
    messhall_uid varchar (255),
    ingredient_uid varchar (255), -- primary key not null default nextval('ingredient_seq'),
    amount integer
);

create table recipe_ingredients (
    ingredient_uid varchar (255), -- primary key not null default nextval('ingredient_seq'),
    recipe_uid varchar (255),
    amount integer
);

create table ingredient (
    ingredient_uid varchar (255) primary key not null default nextval('ingredient_seq'),
    ingredient_name varchar(255),
    calories integer
    -- constraint fk_ingredients foreign key (ingredient_uid) references recipe_ingredients (ingredient_uid)
);