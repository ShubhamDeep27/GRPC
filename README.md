
# GRPC 

 GRPC API for User Login


You need the following library to run:

`go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`

`go get -u google.golang.org/grpc`

`protoc --go_out=plugins=grpc:. *.proto`

## Usage/Examples

```javascript
$ git clone https://github.com/int1359/Grpc.git
[NAME APP]
```


## Try it Out

To compile and run the server

```bash
  $ go run server/main.go

```
Likewise, to run the client:

```bash
  $ go run client/main.go

```
## Tech

API Server technology stack is

* Server code: Golang 1.19
* GRPC (Server and Client Side)
* Database: MySQL
* ORM: gorm 
