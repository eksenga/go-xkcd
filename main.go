package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/emil-io/go-xkcd/client"
)

func main() {
	comicNo := flag.Int(
		"n", int(client.LatestComic), "Comic number to fetch (defaults to latest)",
	)

	clientTimeout := flag.Int64(
		"t", int64(client.DefaultLatestClientTimeout.Seconds()), "Client timeout in seconds",
	)

	saveImage := flag.Bool(
		"s", false, "Save image to current directory",
	)

	outputType := flag.String(
		"t", text, "Print output in format: text/json",
	)

	flag.Parse()

	// instantiate
	xkcdClient := client.NewXKCDClient()
	xkcdClient.SetTimeout(time.Duraction(*clientTimeout) * time.Second)

	// Fetch from API
	comic, err := xkcdClient.Fetch(client.ComicNumber(*comicNo), *saveImage)
	if err != nil {
		log.Println(err)
	}

	if *outputType == "json" {
		fmt.Println(comic.JSON())
	} else {
		fmt.Println(comic.PrettyString())
	}
}
