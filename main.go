package main

import (
	"fmt"
	"github.com/zhukowladimir/qr-interface-for-messenger/qr"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/skip2/go-qrcode"
)

func main() {
	qrEncoder := qr.NewEncoder(qrcode.Medium)
	err := qrEncoder.EncodeText("bababuya", "./test-qr.png")
	if err != nil {
		panic(fmt.Sprintf("encoding failed %s", err.Error()))
	}

	qrScanner := qr.NewScanner()
	results, err := qrScanner.ScanFromImage("./test-qr.png")
	if err != nil {
		panic(fmt.Sprintf("decoding failed %s", err.Error()))
	}

	for _, result := range results {
		fmt.Println(result.GetText())
	}
}
