#1

>docker pull mysql:8

>docker images

    REPOSITORY   TAG       IMAGE ID       CREATED       SIZE
    mysql        8         b05128b000dd   12 days ago   516MB

>mkdir /opt/mysql/backup

>docker run --rm --name mysql-docker -v /opt/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=mysql -p 3306:3306 -d  mysql:8

>docker ps

    CONTAINER ID   IMAGE     COMMAND                  CREATED              STATUS              PORTS                                                  NAMES
    61c759c59e76   mysql:8   "docker-entrypoint.s…"   About a minute ago   Up About a minute   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   mysql-docker

>docker exec -it mysql-docker /bin/bash

>cd /var/lib/mysql/backup/

>mysql -u root -p

mysql> SHOW DATABASES;
    +--------------------+
    | Database           |
    +--------------------+
    | information_schema |
    | mysql              |
    | performance_schema |
    | sys                |
    +--------------------+
    4 rows in set (0.00 sec)

mysql> \h

    For information about MySQL products and services, visit:
       http://www.mysql.com/
    For developer information, including the MySQL Reference Manual, visit:
       http://dev.mysql.com/
    To buy MySQL Enterprise support, training, or other products, visit:
       https://shop.mysql.com/
    
    List of all MySQL commands:
    Note that all text commands must be first on line and end with ';'
    ?         (\?) Synonym for `help'.
    clear     (\c) Clear the current input statement.
    connect   (\r) Reconnect to the server. Optional arguments are db and host.
    delimiter (\d) Set statement delimiter.
    edit      (\e) Edit command with $EDITOR.
    ego       (\G) Send command to mysql server, display result vertically.
    exit      (\q) Exit mysql. Same as quit.
    go        (\g) Send command to mysql server.
    help      (\h) Display this help.
    nopager   (\n) Disable pager, print to stdout.
    notee     (\t) Don't write into outfile.
    pager     (\P) Set PAGER [to_pager]. Print the query results via PAGER.
    print     (\p) Print current command.
    prompt    (\R) Change your mysql prompt.
    quit      (\q) Quit mysql.
    rehash    (\#) Rebuild completion hash.
    source    (\.) Execute an SQL script file. Takes a file name as an argument.
    status    (\s) Get status information from the server.
    system    (\!) Execute a system shell command.
    tee       (\T) Set outfile [to_outfile]. Append everything into given outfile.
    use       (\u) Use another database. Takes database name as argument.
    charset   (\C) Switch to another charset. Might be needed for processing binlog with multi-byte                     charsets.
    warnings  (\W) Show warnings after every statement.
    nowarning (\w) Don't show warnings after every statement.
    resetconnection(\x) Clean session context.
    query_attributes Sets string parameters (name1 value1 name2 value2 ...) for the next query to p                    ick up.
    
    For server side help, type 'help contents'

mysql> \s

    --------------
    mysql  Ver 8.0.27 for Linux on x86_64 (MySQL Community Server - GPL)
    
    Connection id:          11
    Current database:
    Current user:           root@localhost
    SSL:                    Not in use
    Current pager:          stdout
    Using outfile:          ''
    Using delimiter:        ;
    Server version:         8.0.27 MySQL Community Server - GPL
    Protocol version:       10
    Connection:             Localhost via UNIX socket
    Server characterset:    utf8mb4
    Db     characterset:    utf8mb4
    Client characterset:    latin1
    Conn.  characterset:    latin1
    UNIX socket:            /var/run/mysqld/mysqld.sock
    Binary data as:         Hexadecimal
    Uptime:                 37 min 23 sec
    
    Threads: 2  Questions: 9  Slow queries: 0  Opens: 135  Flush tables: 3  Open tables: 54  Queries per second avg: 0.004

mysql> create database test_db;
 
mysql>use test_db;

mysql> source test_dump.sql;

mysql> select count(*) from orders where price>300;

    +----------+
    | count(*) |
    +----------+
    |        1 |
    +----------+
    1 row in set (0.00 sec)

#2

mysql> CREATE USER "test@localhost"IDENTIFIED WITH mysql_native_password BY "test-pass";
Query OK, 0 rows affected (0.00 sec)

mysql> CREATE USER "test"
    -> IDENTIFIED WITH mysql_native_password BY "test-pass"
    -> WITH MAX_QUERIES_PER_HOUR 100
    -> PASSWORD EXPIRE INTERVAL 180 DAY
    -> FAILED_LOGIN_ATTEMPTS 3 PASSWORD_LOCK_TIME 1
    -> ATTRIBUTE '{"fname": "James", "lname": "James"}';
Query OK, 0 rows affected (0.01 sec)

mysql> GRANT SELECT ON *.* TO  "test";
Query OK, 0 rows affected (0.00 sec)

mysql> select * from information_schema.USER_ATTRIBUTES where user="test";

    +------+------+--------------------------------------+
    | USER | HOST | ATTRIBUTE                            |
    +------+------+--------------------------------------+
    | test | %    | {"fname": "James", "lname": "James"} |
    +------+------+--------------------------------------+
    1 row in set (0.00 sec)

#3

mysql> SET profiling = 1;
Query OK, 0 rows affected, 1 warning (0.00 sec)

mysql> SELECT TABLE_NAME,ENGINE,ROW_FORMAT,TABLE_ROWS,DATA_LENGTH,INDEX_LENGTH FROM information_schema.TABLES WHERE table_name = 'orders' and  TABLE_SCHEMA = 'test_db';

    +------------+--------+------------+------------+-------------+--------------+
    | TABLE_NAME | ENGINE | ROW_FORMAT | TABLE_ROWS | DATA_LENGTH | INDEX_LENGTH |
    +------------+--------+------------+------------+-------------+--------------+
    | orders     | InnoDB | Dynamic    |          5 |       16384 |            0 |
    +------------+--------+------------+------------+-------------+--------------+
    1 row in set (0.00 sec)

mysql>  SHOW PROFILES;

    +----------+------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    | Query_ID | Duration   | Query                                                                                                                                                            |
    +----------+------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    |        1 | 0.00186800 | SELECT TABLE_NAME,ENGINE,ROW_FORMAT,TABLE_ROWS,DATA_LENGTH,INDEX_LENGTH FROM information_schema.TABLES WHERE table_name = 'orders' and  TABLE_SCHEMA = 'test_db' |
    +----------+------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    1 row in set, 1 warning (0.00 sec)

mysql> ALTER TABLE orders ENGINE=MyISAM;
Query OK, 5 rows affected (0.11 sec)
Records: 5  Duplicates: 0  Warnings: 0

mysql>  SHOW PROFILES;

    +----------+------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    | Query_ID | Duration   | Query                                                                                                                                                            |
    +----------+------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    |        1 | 0.00186800 | SELECT TABLE_NAME,ENGINE,ROW_FORMAT,TABLE_ROWS,DATA_LENGTH,INDEX_LENGTH FROM information_schema.TABLES WHERE table_name = 'orders' and  TABLE_SCHEMA = 'test_db' |
    |        2 | 0.10533525 | ALTER TABLE orders ENGINE=MyISAM                                                                                                                                 |
    +----------+------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    2 rows in set, 1 warning (0.00 sec)

mysql> ALTER TABLE orders ENGINE=InnoDB;
Query OK, 5 rows affected (0.02 sec)
Records: 5  Duplicates: 0  Warnings: 0

mysql>  SHOW PROFILES;

    +----------+------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    | Query_ID | Duration   | Query                                                                                                                                                            |
    +----------+------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    |        1 | 0.00186800 | SELECT TABLE_NAME,ENGINE,ROW_FORMAT,TABLE_ROWS,DATA_LENGTH,INDEX_LENGTH FROM information_schema.TABLES WHERE table_name = 'orders' and  TABLE_SCHEMA = 'test_db' |
    |        2 | 0.10533525 | ALTER TABLE orders ENGINE=MyISAM                                                                                                                                 |
    |        3 | 0.02962250 | ALTER TABLE orders ENGINE=InnoDB                                                                                                                                 |
    +----------+------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    3 rows in set, 1 warning (0.00 sec)

#4

> cd /etc/mysql

>vim my.cnf

    [mysqld]
    pid-file        = /var/run/mysqld/mysqld.pid
    socket          = /var/run/mysqld/mysqld.sock
    datadir         = /var/lib/mysql
    secure-file-priv= NULL
    
    #Set IO Speed
    # 0 - скорость
    # 1 - сохранность(установлен по умолчанию)
    # 2 - универсальный параметр
    innodb_flush_log_at_trx_commit=0
    
    #Antelope значение по умолчанию
    #Barracuda  поддерживает компрессию
    innodb_file_format=Barracuda
    
    
    #Размер буфера, который в InnoDB используется для
    #записи информации файлов журналов на диск. Разумный диапазон
    #значений составляет от 1М до 8М. Большой буфер журналов позволяет
    #осуществлять объемные транзакции без записи журнала на диск до
    #завершения транзакции. Поэтому если ваши транзакции отличаются
    #значительными объемами, увеличение буфера журналов сократит
    #количество операций ввода/вывода диска.
    innodb_log_buffer_size=1M
    
    #key_buffer_size — крайне важная настройка при использовании MyISAM-таблиц
    #innodb_buffer_pool_size — не менее важная настройка, но уже для InnoDB.
    innodb_buffer_pool_size=640M
    
    #параметр определяет размер файла бинлога
    max_binlog_size = 100M
    
    # Custom config should go here
    !includedir /etc/mysql/conf.d/
