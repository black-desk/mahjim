//go:generate statik -src=./assets/logo -p logo -ns logo
//go:generate statik -src=./assets/files -p files -ns files

package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/black-desk/mahjim/logo"
	"github.com/black-desk/mahjim/merger"
	"github.com/black-desk/mahjim/parser"
	"github.com/patrickmn/go-cache"
	"github.com/rakyll/statik/fs"
)

var port = flag.Uint("p", 8080, "the port server listen at")
var logo image.Image
var c *cache.Cache

func init() {
	// get logo
	FS, err := fs.NewWithNamespace("logo")
	file, err := FS.Open("/favicon.png")
	defer file.Close()
	if err != nil {
		log.Output(1, err.Error())
	}
	logo, err = png.Decode(file)
	if err != nil {
		log.Output(1, err.Error())
	}
	c = cache.New(time.Duration(20*60*1e9), time.Duration(20*60*1e9))
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	fmt.Println(http.ListenAndServe(":"+strconv.FormatUint(uint64(*port), 10), nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {

	if len(request.URL.Path) <= 1 {
		welcome(writer)
		return
	}

	maj_string := request.URL.Path[1:]

	var img image.Image

	if maj_string == "favicon.ico" {
		img = logo
	} else {
		maj_style_config := request.URL.Query()
		p := parser.GetParser(&maj_string, &maj_style_config)
		key := p.Str()
		if object, got := c.Get(key); got == true {
			img = object.(image.Image)
		} else {
			imgs, err := p.Parse()
			if err != nil {
				writeErr(writer, err)
				return
			}
			img = merger.Merge(imgs)
			c.Add(key, img, 0)
		}
	}
	writeImg(writer, img)
}

func welcome(writer http.ResponseWriter) {
	writer.Write([]byte("<html><body>Welcome to Mahjim, a tool use to generate mahjong images!<br> Docs <a href=\"https://github.com/black-desk/mahjim\">here</a></body></html>"))
}

func writeImg(writer http.ResponseWriter, img image.Image) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		log.Println("unable to encode image.")
	}
	writer.Header().Set("Content-Type", "image/png")
	writer.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := writer.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func writeErr(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("error : " + err.Error()))
}
