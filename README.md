# go-rgba4444

go-rgba4444 represents RGBA4444 format images.
RGBA4444 format is often used for textures of [Unity](http://unity3d.com/).

The following code is simple RGBA4444 converter.

``` go
package main

import (
	"flag"
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/shogo82148/go-rgba4444"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Println("rgba4444 input output")
		return
	}
	input := flag.Arg(0)
	output := flag.Arg(1)

	reader, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	newImg := rgba4444.New(img.Bounds())
	draw.FloydSteinberg.Draw(newImg, img.Bounds(), img, image.ZP)

	err = png.Encode(f, newImg)
	if err != nil {
		log.Fatal(err)
	}
}
```

## LICENSE

This software is released under the MIT License, see LICENSE.txt.

## godoc

See [godoc](https://godoc.org/github.com/shogo82148/go-rgba4444) for more imformation.
