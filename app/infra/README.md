### reader/
The reader/ package contains implementations of processes dedicated to reading data from the database.

### writer/
The writer/ package is responsible for implementing processes dedicated to writing data to the database.
In some cases, it may also internally utilize the reader operations.
The function names within this package may include a mix of business logic and database operations, considering the nature of the tasks performed.

### cache.go
The cache.go file contains the implementation of initialization logic for Redis or similar caching systems.

### db.go
The db.go file contains the implementation of initialization logic for MySQL or similar database systems.
