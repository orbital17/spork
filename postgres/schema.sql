
drop table users;
CREATE TABLE users (
  id bigserial primary key,
  name varchar(40) NOT NULL,
  email varchar(40) NOT NULL UNIQUE,
  password varchar(60) NOT NULL,
  date_added timestamp default NULL,
  data JSON
);

CREATE TABLE posts (
  id bigserial primary key,
  date_added timestamp default NULL,
  data JSON,
  body text
);