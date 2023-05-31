### Basic Structure
```
hoge/
  xx.go
  xx_test.go
  input.go
  output.go
  usecase.go
```

### xx.go
This file is used to execute business logic such as DBIO based on Input information received from the handler package and return it to the handler as output.
The file name varies depending on the business logic. get.go for acquisition systems, post.go for update systems, and so on.

### xx_test.go
This is the test file for xx.go. The DBIO part is mocked up and tested.

### input.go
This file defines the logic that converts the request information received by the handler for appropriate use within usecase, as well as the structure that holds the converted information.

### output.go
This file defines a structure with only the information necessary to return to the handler the results of the processing performed by usecase.
It is usually used as a JSON structure

### usecase.go
This file defines a structure for organizing readers, writers, etc. used within usecase, integrating the necessary resources within usecase, and providing an interface for executing business logic.
