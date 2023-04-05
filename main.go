package main

import (
	"flag"
	"fmt"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/zhukowladimir/qr-interface-for-messenger/qr"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/skip2/go-qrcode"
)

func encode(text string, qrPath string) {
	qrEncoder := qr.NewEncoder(qrcode.Medium)
	err := qrEncoder.EncodeText(text, qrPath)
	if err != nil {
		panic(fmt.Sprintf("encoding failed %s", err.Error()))
	}
}

func decode(qrPath string) {
	qrScanner := qr.NewScanner()
	results, err := qrScanner.ScanFromImage(qrPath)
	if err != nil {
		panic(fmt.Sprintf("decoding failed %s", err.Error()))
	}

	for _, result := range results {
		fmt.Println(result.GetText())
	}
}

func main() {
	mode := flag.String("mode", "encode", "Program operation mode. 'encode', 'decode'")
	text := flag.String("text", "text should be encoded", "Еhe text that will be encoded into a qr image")
	qrPath := flag.String("qr-path", "./test-qr.png", "Path to qr image")
	flag.Parse()

	if *mode == "encode" {
		encode(*text, *qrPath)
	} else if *mode == "decode" {
		decode(*qrPath)
	} else {
		panic(fmt.Sprintf("Incorrect 'mode' arg value: %s", *mode))
	}
}
