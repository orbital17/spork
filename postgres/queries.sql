INSERT INTO
  users (id, name, password, date_added)
VALUES
  (
    2,
    'John',
    'sample_password',
    current_timestamp
  );


alter table
  users
add
  column email varchar(40);


drop table users;

select * from users;

select
		id, name, email, password
	from
		users
  where email = 'olexiy.tkachenko@gmail.com'

select * from posts;