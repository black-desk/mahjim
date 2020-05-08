package main

import (
	"bytes"
	"fmt"
	"github.com/black-desk/mahjim/Merger"
	"github.com/black-desk/mahjim/Parser"
	"image"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func parse(majs string) (*[]*image.Image, error) {
	parser := Parser.NewParser(strings.Split(majs, "|"))
	err := parser.Parse()
	return parser.Imgs, err
}

func genMajImage(majs string) (*image.Image, error) {
	imgs, err := parse(majs)
	if err == nil {
		img := Merger.Merge(imgs)
		return &img, nil
	} else {
		return nil, err
	}
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
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		majs, _ := url.PathUnescape(request.URL.String()[1:])
		img, err := genMajImage(majs)
		if err == nil {
			writeImg(writer, img)
		} else {
			writeErr(writer, err)
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
