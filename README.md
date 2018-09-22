# Only have static and proxy

## Golang Install
```
go get github.com/ymzuiku/gonginx
```

## Start

start server and load ./static/*

```sh
$ cd project
$ gonginx ./static 4001
```

start server and load ./static/*, and proxy to http://localhost:5000

```sh
$ gonginx ./static 4001 5000
```

## If you no have golang, you can downdown release

osx:

```
curl -o /usr/local/bin/gonginx https://raw.githubusercontent.com/ymzuiku/gonginx/master/bin/osx/gonginx 
chmod 0700 /usr/local/bin/gonginx
```

linux:

```
wget -c -O /usr/local/bin/gonginx https://raw.githubusercontent.com/ymzuiku/gonginx/master/bin/linux/gonginx 
chmod 0700 /usr/local/bin/gonginx
```
and use:

```sh
gonginx ./static 4001
```