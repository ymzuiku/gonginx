# Only have static and proxy

## Release Downdown

osx:

```
curl -o /usr/local/bin/gonginx https://raw.githubusercontent.com/ymzuiku/gonginx/master/bin/osx/gonginx 
chmod 0700 /usr/local/bin/gonginx
```

linux:

```
wget -c -O /usr/local/bin/gonginx https://raw.githubusercontent.com/ymzuiku/gonginx/master/bin/linux/gonginx 
```

## Golang Install
```
go get github.com/ymzuiku/gonginx
```

## Start

start server and load ./static/*

```
cd project
gonginx ./static 4001
```

start server and load ./static/*, and proxy to http://localhost:5000

```
gonginx ./static 4001 5000
```