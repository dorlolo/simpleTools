create table sys_dict_area
(
    pid         integer,
    short_name  text,
    merger_name text,
    level       integer,
    pinyin      text,
    phone_code  text,
    zip_code    text,
    first       text,
    lng         text,
    lat         text,
    area_code   text,
    name        text
);

create unique index sys_dict_area_bak_id_uindex
    on sys_dict_area_bak (id);