# Preparations

None. Just normal go stuff

# Write the code
Configure an run openapi as documented here https://github.com/oapi-codegen/oapi-codegen
Result: [gen.go](api/gen.go)

## Connect generated code with framework
Implement the generated Interface `ServerInterface` see [impl.go](api/impl.go)
Instantiate the implementation and use the generated factory function `RegisterHandlersWithBaseURL`
in main.go

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

curl -X POST -H "Content-Type: application/json" -d '{"name":"cpil_posted"}' "localhost:8080/helloworld"
// {"message":"Hello, cpil_posted!"}
```

## Current status:
I works nicely. The abstraction is good. Generation of server code seems to be good. The oapi-codegen tool
has more features. Not tested: Generation of client code. Not implemented: Service a "playable" swagger/openapi
html interface.

In general this seems like a good and clean solution for a contract first API architecture.

## Notes
There are a lot of alternatives for contract first codegen solutions e.g.
https://github.com/OpenAPITools/openapi-generator with more than 21k gitshub stars and more than 3k contributors.
