wordcolor.go
------------

[![Build Status](https://travis-ci.org/lichunqiang/wordcolor.go.svg?branch=master)](https://travis-ci.org/lichunqiang/wordcolor.go)

Color your words! Similar words have Similar color.

## Requirements

Go 1.2 or above.

## Install & Usage

```
$ go get github.com/lichunqiang/wordcolor.go
```

```
import "github.com/lichunqiang/wordcolor.go"

func main () {
    println(WordColor("Hello world")) //rgb(255, 56, 13)
}
```

## License

MIT