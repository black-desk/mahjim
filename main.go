package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/black-desk/mahjim/merger"
	"github.com/black-desk/mahjim/parser"
)

var logo image.Image

func init() {
	// get logo
	file, err := os.Open("./favicon.png")
	defer file.Close()
	if err != nil {
		log.Output(1, err.Error())
	}
	logo, err = png.Decode(file)
	if err != nil {
		log.Output(1, err.Error())
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println(http.ListenAndServe(":8080", nil)) //TODO: port config
}

func handler(writer http.ResponseWriter, request *http.Request) {

	maj_string := request.URL.Path[1:] // TODO: handle favicon.ico maj_config := request.URL.Query()

	var img image.Image

	if maj_string == "favicon.ico" {
		img = logo
	} else {
		maj_style_config := request.URL.Query()
		p := parser.GetParser(&maj_string, &maj_style_config)
		imgs, err := p.Parse()
		if err != nil {
			writeErr(writer, err)
		}
		img = merger.Merge(imgs)
	}
	writeImg(writer, &img)
}

func writeImg(writer http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}
	writer.Header().Set("Content-Type", "image/png")
	writer.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := writer.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func writeErr(writer http.ResponseWriter, err error) {
	header := writer.Header()
	header.Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(writer, err)
}
