create table user (
    username VARCHAR(100) NOT NULL,
    password varchar(100) NOT NULL,
    primary key (username)
) ENGINE = InnoDB;
insert into user(username, password)
values('admin', 'admin');