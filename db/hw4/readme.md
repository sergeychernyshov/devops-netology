#1

>docker pull postgres:13

>mkdir /opt/postgres13/backup

>mkdir /opt/postgres13/data

>docker run --rm --name pg-docker -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 -v /opt/postgres13/data/:/var/lib/postgresql/data -v /opt/postgres13/backup:/var/lib/postgresql/backup postgres:13

>docker ps

    CONTAINER ID   IMAGE         COMMAND                  CREATED          STATUS          PORTS                                       NAMES
    e91eaf9a1023   postgres:13   "docker-entrypoint.s…"   12 seconds ago   Up 11 seconds   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   pg-docker

>docker exec -it pg-docker /bin/bash

>su postgres

>psql

###Вывод списка БД
    postgres=# \l+
                                                                       List of databases
       Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges   |  Size   | Tablespace |                Description
    -----------+----------+----------+------------+------------+-----------------------+---------+------------+--------------------------------------------
     postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 |                       | 7901 kB | pg_default | default administrative connection database
     template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +| 7753 kB | pg_default | unmodifiable empty database
               |          |          |            |            | postgres=CTc/postgres |         |            |
     template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +| 7753 kB | pg_default | default template for new databases
               |          |          |            |            | postgres=CTc/postgres |         |            |
    (3 rows)

###Подключения к БД
    postgres-# \c postgres
    You are now connected to database "postgres" as user "postgres".

###Вывода списка таблиц
    postgres-# \dt
    Did not find any relations.

Просматриваю служебные объекты используя параметр S

    postgres-# \dtS
                    List of relations
       Schema   |          Name           | Type  |  Owner
    ------------+-------------------------+-------+----------
     pg_catalog | pg_aggregate            | table | postgres
    ............
     pg_catalog | pg_user_mapping         | table | postgres
    (62 rows)

###Вывода описания содержимого таблиц
    postgres-# \dS pg_tablespace
                                    Table "pg_catalog.pg_tablespace"
       Column   |   Type    | Collation | Nullable | Default | Storage  | Stats target | Description
    ------------+-----------+-----------+----------+---------+----------+--------------+-------------
     oid        | oid       |           | not null |         | plain    |              |
     spcname    | name      |           | not null |         | plain    |              |
     spcowner   | oid       |           | not null |         | plain    |              |
     spcacl     | aclitem[] |           |          |         | extended |              |
     spcoptions | text[]    | C         |          |         | extended |              |
    Indexes:
        "pg_tablespace_oid_index" UNIQUE, btree (oid), tablespace "pg_global"
        "pg_tablespace_spcname_index" UNIQUE, btree (spcname), tablespace "pg_global"
    Tablespace: "pg_global"
    Access method: heap

###Выход из psql
    postgres-# \q
    postgres@e91eaf9a1023:/$


#2
  
>psql

    postgres=# CREATE DATABASE test_database;
    CREATE DATABASE

    postgres=# \q
    postgres@e91eaf9a1023:/$

>cd /var/lib/postgresql/backup

>psql -U postgres -f test_dump.sql test_database
    SET
    SET
    SET
    SET
    SET
     set_config
    ------------
    
    (1 row)
    
    SET
    SET
    SET
    SET
    SET
    SET
    CREATE TABLE
    ALTER TABLE
    CREATE SEQUENCE
    ALTER TABLE
    ALTER SEQUENCE
    ALTER TABLE
    COPY 8
     setval
    --------
          8
    (1 row)
    
    ALTER TABLE

>psql

    psql (13.5 (Debian 13.5-1.pgdg110+1))
    Type "help" for help.

    postgres=# \c test_database
    You are now connected to database "test_database" as user "postgres".

    test_database=# \dt
         List of relations
     Schema |  Name  | Type  |  Owner
    --------+--------+-------+----------
     public | orders | table | postgres
    (1 row)

    test_database=# ANALYZE VERBOSE public.orders;
    INFO:  analyzing "public.orders"
    INFO:  "orders": scanned 1 of 1 pages, containing 8 live rows and 0 dead rows; 8 rows in sample, 8 estimated total rows
    ANALYZE

    test_database=# select max(avg_width) from pg_stats where tablename='orders';
     max
    -----
      16
    (1 row)

#3

    test_database=# \d public.orders_backup
                                Table "public.orders_backup"
     Column |         Type          | Collation | Nullable |              Default
    --------+-----------------------+-----------+----------+------------------------------------
     id     | integer               |           | not null | nextval('orders_id_seq'::regclass)
     title  | character varying(80) |           | not null |
     price  | integer               |           |          | 0
    Indexes:
        "orders_pkey" PRIMARY KEY, btree (id)

    test_database=# create table orders (id integer, title varchar(80), price integer) partition by range(price);
    CREATE TABLE
    test_database=# create table orders_less499 partition of orders for values from (0) to (499);
    CREATE TABLE
    test_database=# create table orders_more499 partition of orders for values from (499) to (999999999);
    CREATE TABLE
    test_database=# begin transaction;
    BEGIN
    test_database=# insert into orders (id, title, price) select * from orders_backup;
    INSERT 0 8
    test_database=# commit;
    COMMIT
    test_database=# ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;
    ALTER SEQUENCE
    ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);
    ALTER TABLE
    test_database=# ALTER TABLE ONLY public.orders_less499 ADD CONSTRAINT orders_pkey_less499 PRIMARY KEY (id);
    ALTER TABLE
    test_database=# ALTER TABLE ONLY public.orders_more499  ADD CONSTRAINT orders_pkey_more499  PRIMARY KEY (id);
    ALTER TABLE
    test_database=# alter table public.orders_less499 alter column id set not null;
    ALTER TABLE
    test_database=# alter table public.orders_more499 alter column id set not null;
    ALTER TABLE
    test_database=# alter table public.orders_less499 alter title set not null;
    ALTER TABLE
    test_database=# alter table public.orders_more499 alter column title set not null;
    ALTER TABLE
    test_database=# ALTER TABLE ONLY public.orders_less499  ALTER COLUMN price SET DEFAULT 0;
    ALTER TABLE
    test_database=# ALTER TABLE ONLY public.orders_more499  ALTER COLUMN price SET DEFAULT 0;
    ALTER TABLE

    test_database=# alter table public.orders alter column id set not null;
    ALTER TABLE
    test_database=# alter table public.orders alter column title set not null;
    ALTER TABLE
    test_database=# alter table public.orders ALTER COLUMN price SET DEFAULT 0;
    ALTER TABLE
    test_database=# ALTER TABLE ONLY public.orders_more499 ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);
    ALTER TABLE
    test_database=# ALTER TABLE ONLY public.orders_less499 ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);
    ALTER TABLE



    test_database=# \d public.orders
                                 Partitioned table "public.orders"
     Column |         Type          | Collation | Nullable |              Default
    --------+-----------------------+-----------+----------+------------------------------------
     id     | integer               |           | not null | nextval('orders_id_seq'::regclass)
     title  | character varying(80) |           | not null |
     price  | integer               |           |          | 0
    
    test_database=# \d public.orders_more499
                                   Table "public.orders_more499"
     Column |         Type          | Collation | Nullable |              Default
    --------+-----------------------+-----------+----------+------------------------------------
     id     | integer               |           | not null | nextval('orders_id_seq'::regclass)
     title  | character varying(80) |           | not null |
     price  | integer               |           |          | 0
    Partition of: orders FOR VALUES FROM (499) TO (999999999)
    Indexes:
        "orders_pkey_more499" PRIMARY KEY, btree (id)
    
    test_database=# \d public.orders_less499
                                   Table "public.orders_less499"
     Column |         Type          | Collation | Nullable |              Default
    --------+-----------------------+-----------+----------+------------------------------------
     id     | integer               |           | not null | nextval('orders_id_seq'::regclass)
     title  | character varying(80) |           | not null |
     price  | integer               |           |          | 0
    Partition of: orders FOR VALUES FROM (0) TO (499)
    Indexes:
        "orders_pkey_less499" PRIMARY KEY, btree (id)

    test_database=# drop table orders_backup;
    DROP TABLE

Можно исключить "ручное" разбиение при проектировании таблицы orders, если изначально сделать таблицу секционированной.

#4

    test_database=# \q
    postgres@e91eaf9a1023:~/backup$

>pg_dump -U postgres -d test_database >test_database_dump_db.sql

Можно добавить уникальный индекс 
CREATE INDEX ON orders (title);
или если нужно обеспечить уникальность не зависимо от регистра
CREATE INDEX ON orders (lower(title)); 




