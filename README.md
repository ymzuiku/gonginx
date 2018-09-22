# Only have static and proxy

## Install
```
go get github.com/ymzuiku/gonginx
```

## Start

start server and load ./static/*

```
cd project
gonginx ./static
```

start server and load ./static/*, and proxy to http://localhost:5000

```
gonginx ./static 5000
```