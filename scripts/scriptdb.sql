/*drop database if exists testDB;

create database testDB;
*/

 CREATE TABLE authors (
  id serial PRIMARY KEY,
  name text,
  date_of_birth timestamp
);
                        
CREATE TABLE books (
  id serial PRIMARY KEY,
  author_id integer REFERENCES authors (id),
  isbn text
);
                        
CREATE TABLE sale_items (
  id serial PRIMARY KEY,
  book_id integer REFERENCES books (id),
  customer_name text,
  item_price money,            
  quantity integer
);    

insert into authors(name,date_of_birth) values ('jainor',current_timestamp);
insert into authors(name,date_of_birth) values ('nestor',current_timestamp);
insert into authors(name,date_of_birth) values ('Lorelai Gilmore',current_timestamp);
insert into authors(name,date_of_birth) values ('Lorel',current_timestamp);


insert into books(author_id,isbn) values (1,'123-123-123');
insert into books(author_id,isbn) values (1,'323-123-123');
insert into books(author_id,isbn) values (2,'125-123-123');
insert into books(author_id,isbn) values (2,'327-123-123');
insert into books(author_id,isbn) values (3,'154-123-123');
insert into books(author_id,isbn) values (3,'723-123-123');
insert into books(author_id,isbn) values (4,'345-567-353');
insert into books(author_id,isbn) values (4,'567-345-134');

insert into sale_items(book_id,customer_name,item_price,quantity) values (1,'buyer0',123.00,1);
insert into sale_items(book_id,customer_name,item_price,quantity) values (1,'buyer1',23.00,2);
insert into sale_items(book_id,customer_name,item_price,quantity) values (3,'buyer2',65.00,4);
insert into sale_items(book_id,customer_name,item_price,quantity) values (5,'buyer3',70.00,2);
insert into sale_items(book_id,customer_name,item_price,quantity) values (7,'buyer4',94.00,3);
