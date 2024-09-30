# Preparations

None. Just normal go stuff

# Write the code
Write the server with annotations

## Swagger Annotations
* @Summary: A brief explanation of what the endpoint does (e.g., "getHelloWorld returns hello world").
* @Description: A detailed description of the endpoint’s functionality.
* @Tags: Used to categorize endpoints (in this case, all endpoints are under helloworld).
* @Param: Defines request parameters, their location (query, body), type, required state, and examples.
* @Success and @Failure: Define possible response codes, their descriptions, and the schema of the response body.
* @Router: Defines the HTTP method and the route for the endpoint.



## Generate swagger file from annotated code
* run `swag init` to create the documentation based on the annotations. Will be in `docs` per default

# Build and run
Just like any other go project

# check if it works

```shell
curl localhost:8080/helloworld
// {"message":"Hello, World!"}
```


```shell
curl localhost:8080/helloworld
// {"message":"Hello, World!"}

curl -X PUT "localhost:8080/helloworld?name=cpil"
// {"message":"Hello, cpil!"}


curl -X POST -H "Content-Type: application/json" -d '{"name":"cpil_posted"}' "localhost:8080/helloworld"
// {"message":"Hello, cpil_posted!"}
```

