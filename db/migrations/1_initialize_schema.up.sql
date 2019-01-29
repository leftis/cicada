CREATE TABLE administrators (
   id integer,
   username text,
   hashed_pwd text,
   last_login_at timestamp,
   last_login_ip cidr,
   PRIMARY KEY( id )
);