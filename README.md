# counter-demo

A demo web server that counts lines of source code in a given Go package
```
go get github.com/rosenhouse/counter-demo
cd $GOPATH/src/github.com/rosenhouse/counter-demo
go run server/main.go &
curl localhost:8000/lines/github.com/rosenhouse/counter-demo
kill %1
```

## Huh?
This is the material I used in my talk on TDDing with Mocks in Go.

Follow along via [the commit history](https://github.com/rosenhouse/counter-demo/commits/master).
