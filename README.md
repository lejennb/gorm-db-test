# Test: GORM / MariaDB Container
* Use an `init.sql` script?
```sql
CREATE DATABASE IF NOT EXISTS test;

USE test;

...
```

* Have GORM handle the table creation or move all that to the script too?
```sql
-- Examples
CREATE TABLE IF NOT EXISTS authors (
    id INT PRIMARY KEY,
    name TEXT
);

CREATE TABLE IF NOT EXISTS books (
    id INT PRIMARY KEY,
    title TEXT,
    author INT, FOREIGN KEY(author) REFERENCES author(id) ON DELETE CASCADE
);

...
```
* Gorm will map table names of the structs to plural forms of the name in snake case!

* Run (examples):
```bash
docker-compose up

docker exec -it mariadb01 bash
docker exec -it mariadb01 mysql --user=root -p test
```

