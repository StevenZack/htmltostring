# htmltostring: convert HTML/JS/CSS files into Go string

# Installation
```go 
go get github.com/StevenZack/htmltostring
```

# Get started

## 1. Create a `hello` directory under your $GOPATH/src , and create a `main.go` file in it.

```go
package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter,r *http.Request){
	
}
```

## 2. Create a `html` directory under `hello`, and create a `index.html` in it.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    Hello world
</body>
</html>
```

## Then the directory structure will be like this:
```shell
hello/
├── html
│   └── index.html
└── main.go
```

## 3.Run `htmltostring` under `hello`
```shell
$ cd $GOPATH/src/hello
$ htmltostring
```

## Then the directory structure goes like this:
```shell
hello/
├── html
│   └── index.html
├── main.go
└── views
    └── index.go
```

  Look , it just created a `index.go` file automatically , based on your `index.html` :
  ```go
package views

var Str_index =`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    Hello world
</body>
</html>`
```

## Now you can use `index.html` as a Go string , in your `main.go`:
```go
package main

import (
	"hello/views"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(views.Str_index))
}
```
