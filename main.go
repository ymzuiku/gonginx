package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/NYTimes/gziphandler"
)

var dir = "./static"

var isDev = runtime.GOOS == "darwin" || runtime.GOOS == "windows"
var port = "4001"
var proxy = "noProxy"

func main() {
	if isDev == false {
		port = "80"
	}
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	if len(os.Args) > 2 {
		port = os.Args[2]
	}
	if len(os.Args) > 3 {
		proxy = os.Args[3]
	}
	withGZ := gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if isDev {
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
		}
		staticHandler := http.FileServer(http.Dir(dir))
		staticHandler.ServeHTTP(w, req)
	}))
	if proxy == "noProxy" {
		http.Handle("/", withGZ)
	} else {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   "127.0.0.1:" + proxy,
		})
		http.Handle("/", middlewarel(withGZ, proxy))
	}

	fmt.Println("static: http://127.0.0.1:" + port)
	server := &http.Server{Addr: ":" + port}
	server.ListenAndServe()
}

func middlewarel(next http.Handler, proxy http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if len(url) == 4 && (url == "/401" || url == "/403") {
			url = "/"
			http.Redirect(w, r, "/", 300)
			return
		}
		re := regexp.MustCompile("(.js|.html|.css|.png|.ico|.map|.woff|.svg|.eot|.ttf|.jpg|.xml|.json|.gif)$")

		if re.Find([]byte(url)) != nil || url == "/" || strings.Contains(r.Header.Get("Accept"), "text/html") {
			next.ServeHTTP(w, r)
		} else {
			proxy.ServeHTTP(w, r)
		}
	})
}
