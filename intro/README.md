# Introduction
This is an introduction to protobuf. It covers topics such as scalar types, repeated field, enumerations, comments, 
nested message type, importing proto files, package, any, oneof, map, reading and writing protobuf messages and schema
evolution. 


# Creating the go module
```shell
mkdir intro
cd intro
go mod init github.com/garcios/go-grpc/intro
```
The above commands will create go.mod file in the intro directory with the following content:
```go
module github.com/garcios/go-grpc/intro

go 1.22.2
```


## Requirements

- Protobuf Compiler
```shell
brew install protobuf
```

- Protobuf Go Plugin
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

- Open API
```shell
go get github.com/google/gnostic
go install github.com/google/gnostic/cmd/protoc-gen-openapi
```

- Make
```shell
brew install make
```

## Proto compiler command
```shell
cd src
protoc --go_out=. ./proto/basic/*.proto
```

or using go generate command
```shell
cd <project root directory>
go generate ./...
```

or using make command
```shell
make generate
```

## Update go modules to satisfy golang imports for the generated code
```shell
cd ..
go mod tidy
```

or using make command
```shell
make tidy
```

## Run the code
```shell
cd src
go run .
```

or using make command
```shell
make run
```

## Using protojson
When you're using Protobuf as your data structure, you might want to use protojson for the following reasons:
1. __Native Compatibility__: protojson is built to work with Protobuf. It directly understands Protobuf message 
definitions, enumerations, nested messages, etc. It's more seamless and natural to marshal/unmarshal using protojson 
if your primary data structure is Protobuf.
2. __Efficiency__: protojson is designed to be fast and efficient with Protobuf data structures. The standard 
encoding/json library isn't optimized specifically for Protobuf.
3. __Additional Features__: protojson also provides additional options that are not available in encoding/json, 
such as emitting fields with default values when marshaling, treating null values as empty messages, etc. For example, 
this code block adapts the JSON marshaller settings:

```go
import "google.golang.org/protobuf/encoding/protojson"

marshaler := protojson.MarshalOptions{
	EmitUnpopulated: true,
}
jsonBytes, err := marshaler.Marshal(protoMessage)
if err != nil {
	// handle error
}
```
Remember, if you're not working with Protobuf data, then encoding/json is perfectly suitable for most json 
serialization/deserialization tasks. But when working with Protobuf, protojson could provide a more native and efficient
option.


## Reading and Writing Protobuf to Disk
Basic algorithm to write protobuf message to binary file:
```go
bytes, err := proto.Marshal(protoMessage)
os.WriteFile(filename, bytes, permission)
```

Basic algorith to read binary file into protobuf message:
```go
bytes, err := os.ReadFile(filename)
proto.Unmarshall(bytes, protoMessage)
```

## Schema Evolution
When dealing with Protocol Buffers (protobuf) schema evolution, there are several important topics to consider to 
maintain compatibility. In particular, these principles aim to ensure that old versions of the software can handle new
data structures and that new software versions can handle old data structures correctly. Following these principles will
help to avoid pitfalls of backwards and forward compatibility.

1. __Field Numbers__: Field numbers in your protobuf message definitions are used to identify your fields in the message
binary format, and should never be changed. The protocol buffer language takes care of maintaining the correct field 
numbers in the binary data, allowing you to rename fields as developments evolve without needing to worry about field 
numbers.
   
2. __Adding Fields__: You can add new fields to your message formats without affecting backwards-compatibility. 
Any unrecognized fields will be ignored, so old versions of the software can still read new data formats.

3. __Deleting Fields__: Never delete fields – mark them as deprecated. Protobuf does not release field numbers; 
any fields that you remove should be deprecated and their numbers should not be used in your new message types. If such 
a field number is reused, data corruption and compatibility issues can occur.

4. __Renaming Fields__: Renaming a field does not break either backward compatibility or forward compatibility. 
This is because protobuf uses numeric identifiers (tags or field numbers) rather than field names to identify fields 
in the message binary format. 

5. __Data Types__: You can change a field from a compatible type to another. For example, int32, uint32, int64, uint64, 
and bool are all compatible types.

6. __Repeated Fields__: Keep in mind the changes inside a repeated field. Changing a single item to repeated could cause
issues in terms of compatibility.

7. __Using Reserved__: You can specify reserved fields in your messages. So when you delete a field, you may want to 
reserve that number for the future to prevent others from reusing it in some unrelated manner.

### Forward compatibility
Forward compatibility refers to the ability of a system to accept input intended for a later version of itself. The idea
is that you should be able to use old software systems or tools with new data formats or structures, even if those 
formats or structures didn't exist when the software was originally created. 

In essence, if a system is forward compatible, it exhibits a certain level of future-proofing. It signifies that changes
or enhancements to a technology will not completely render the ability of older systems to function and interact 
correctly.

In the context of Protocol Buffers (protobuf), forward compatibility means that older versions of a program can still 
process data that contains new field types, perhaps by skipping them. They might not fully understand the new data 
fields, but they won't break when encountering them. This is crucial in many environments where different systems and 
components may be using different versions of the software, and it's not always feasible to update all systems concurrently.

For example, let's consider the following scenario to show an example of forward compatibility with Protocol Buffers (protobuf).
Assume you have a version 1 of your software running in production that uses the following protobuf message:
```protobuf
// Message User in version 1
message User {
    int32 id = 1;
    string email = 2;
    string name = 3;
}
```

Now, you're developing version 2 of your software, and you decide to add a new field to your User message like so:
```protobuf
// Message User in version 2
message User {
    int32 id = 1;
    string email = 2;
    string name = 3;
    string phone = 4; // Added field
}
```
The new field phone has been added with a new field number 4.

__Forward Compatibility__: The old version 1 of your software can read and ignore the new field phone from the new version 
of User messages. It will continue to function correctly with its existing functionality, even if it doesn't recognize 
or use the new phone field information.

This is a simple but effective way to illustrate how protobuf maintains forward compatibility. Thanks to this feature, 
you can evolve your data structures over time and always be confident that you won't break your older systems as a result.


### Backward compatibility
Backward compatibility, also known as downward compatibility, refers to the ability of newer systems or technology to 
understand, function, or operate with older inputs or older versions of the same system. In other words, a system is 
backward compatible if newer versions of it can accept the same input and produce the same output as the older versions.

If a piece of software is backward compatible, it can interpret data generated by an older version of itself or interact
with other systems that are older versions of itself without any issues.

In the context of Protocol Buffers (protobuf), backward compatibility means that your newer software versions can 
process older data streams that lack newer features added in the newer versions.

This is very important because it ensures that upgrades to a system don't disrupt the functioning of the system and 
don't necessitate simultaneous updates to all interacting systems. Applications don't have to be migrated and elucidated
at the same time as the system updates, leading to smoother and more efficient development and operation.

For example, Let's assume you have a version 1 of your software running in production that uses the following protobuf 
message:
```protobuf
// Message User in version 1
message User {
    int32 id = 1;
    string email = 2;
    string name = 3;
    string phone = 4; // Existing field
}
```
Then, you're developing a version 2 of your software, and you decide to remove the phone field from your User message 
like so:
```protobuf
// Message User in version 2
message User {
    int32 id = 1;
    string email = 2;
    string name = 3;
    // string phone = 4; // Removed field
}
```

__Backward Compatibility__: The new version 2 of your software can still read the old messages that include the phone 
field. It will simply ignore the phone field because it does not exist in the newer protobuf definition, and the software
continues to work correctly using the fields it recognizes.

This denotation is a simple yet powerful way to demonstrate how protobuf supports backward compatibility. This 
flexibility allows developers to safely evolve their data structures in an interoperable manner over time, reducing the 
risk that updates will disrupt system operation.


## References
- https://github.com/uber/prototool/blob/dev/style/README.md
- https://cloud.google.com/apis/design
- https://github.com/uber-go/guide/blob/master/style.md