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
