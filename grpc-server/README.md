# gRPC server

## Creating go module
```shell
go mod init github.com/garcios/go-grpc/grpc-server
go mod tidy
```

## Hexagonal Architecture
This module uses Hexagonal Architecture. Hexagonal Architecture, also known as Ports and Adapters pattern, is a software
design pattern that proposes how to structure an application such that it remains separate from the outside world. 
The central idea is that the application/core logic should not depend on databases, messaging systems, and frameworks.

Key features of hexagonal architecture include:

1. __Central Domain__: The application/core logic sits at the center (the hexagon), shaping the architecture. This central 
domain contains all business logic and no code related to infrastructure, interfaces, etc.

2. __Ports__: Ports define the boundary of your application. They are interfaces that represent all possible interactions 
with the application. There are two types of ports:
 
 - __Primary or driving ports__: Represent the use cases of your application which are driven by the actors (user, system...).
 - __Secondary or driven ports__: Used by your application to interact with external actors for example, database or external
     service.

The purpose of Hexagonal Architecture is to allow an application to equally be driven by users, programs, automated test
or batch scripts, and to be developed and tested in isolation from its eventual run-time devices and databases.

Here's an example in GoLang of an application that follows the hexagonal architecture:

```go
package main

type Port interface {
   // defines the methods that our application/core logic exposes.
}

type Adapter struct {
   // implements the Port interface and actually contains logic.
}

func main() {
   var primaryAdapter Port = &Adapter{}

   // Application logic can now be driven by primaryAdapter.
}
```
Keep in mind that this is a simplistic representation of Hexagonal Architecture, and real-world implementations may have
various degrees of complexity and additional components depending on the project requirements.