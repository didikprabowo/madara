# Sample Rest API using GIN Framework

a tiny example REST API using [Gin Framework][gin], highly influenced by [this repo][gin-boilerplate]

## How to use

setup your base golang workspace, export your `GOPATH`, then execute:

```go get github.com/widnyana/madara```

create `main.go` and add following code:

```
package main

import (
    "github.com/widnyana/madara"
    _ "github.com/widnyana/madara/endpoint/users"
)

func main() {
    app := madara.New()
    app.Router.Run("127.0.0.1:8080")
}
```

run it with `go run main.go`

if no error, your api is live in `127.0.0.1:8080/users`

## Thanks

- a beautiful yet powerfull [GIN Framework][gin]
- [cornfeedhobo][gin-boilerplate] for his cool boilerplate


## License

provided as-is, for educational purpose, responsibility held by the users

see `LICENSE`


[gin]:https://github.com/gin-gonic/gin
[gin-boilerplate]:https://github.com/cornfeedhobo/gin-boilerplate