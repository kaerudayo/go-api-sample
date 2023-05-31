### handler/
The handler/ package typically includes functions or methods that are responsible for handling specific HTTP endpoints or routes.
These handlers parse request parameters, validate input, invoke the corresponding use case methods, and construct the HTTP response.

### middleware/
The middleware/ package contains the implementation of router middleware.
Middleware is used to perform additional processing or filtering before and after handling a request.
For example, common tasks like authentication can be implemented as middleware.

### router/
The router/ package contains the implementation responsible for routing handlers.
The router is responsible for routing requests to the appropriate handler.
Typically, the appropriate handler is selected based on the path or method of the HTTP request.
