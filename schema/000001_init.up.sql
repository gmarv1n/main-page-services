CREATE TABLE users
(
    id              serial              not null unique,
    name            varchar(255)        not null,
    username        varchar(255)        not null unique,
    password_hash   varchar(255)        not null
);

CREATE TABLE quotes
(
    id                      serial              not null unique,
    quote                   text                not null,
    record_id               int                 not null,
    region_id               int                 not null,
    author                  text,
    author_description      text
);

CREATE TABLE themes
(
    id                      serial              not null unique,
    theme_id                int                 not null,
    region_id               int                 not null
);

CREATE TABLE picture_of_the_day_tabs_lists
(
    id              serial              not null unique,
    region_id       int                 not null
);

CREATE TABLE picture_of_the_day_tabs_items
(
    id              serial                                                not null unique,
    region_id       references picture_of_the_day_tabs_lists (region_id)  not null,
    type            varchar(255)                                          not null,
    entity_id       int                                                   not null,
    order_place     int                                                   not null,
    title           text                                                  not null
);
