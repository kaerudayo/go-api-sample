### entity/
The entity/ package contains structures that represent a one-to-one relationship with database tables.
Each entity corresponds to a specific table and is directly associated with the structure and columns of that table.
Entities are used to represent data from the database table.

### model/
The model/ package includes structures that make it easier to handle information retrieved from the database within the application.
It may also include business logic and operations specific to the application.
Generally, the structures used in use cases or similar components are defined in this package.
The purpose of this package is to transform database information into a format that is more suitable for the application's requirements.

### repository/
The repository/ package abstracts the logic for accessing the database into interfaces.
It follows the CQRS (Command Query Responsibility Segregation) principle by separating commands and queries.
Commands handle operations that modify data, such as creating or updating records, while queries handle operations that retrieve or read data.
This package is responsible for abstracting the interaction with the database and handles tasks related to persistence and retrieval of data.

Each package serves a different role and has specific responsibilities, contributing to the effective design of the application's structure and data handling. They help enhance flexibility and maintainability by assigning different functionalities and responsibilities to each package.

Once you have an implementation of the interface and an implementation that satisfies it, you can run `make mockgen SOURCE="file name"` to automatically generate a mock file
