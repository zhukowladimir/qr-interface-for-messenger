package main

import (
	"fmt"
	"image"
	"log"
	"os"

	// import gif, jpeg, png
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strings"

	// import bmp, tiff, webp
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/multi/qrcode"

	rqrcode "github.com/skip2/go-qrcode"
)

// CreateQRCode creates QRCode
func CreateQRCode() {
	err := rqrcode.WriteFile("i am clown", rqrcode.Medium, 256, "qr.png")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("qr code created!")
}

func ReadQRCode() {
	fi, err := os.Open("qr.png")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	img, _, err := image.Decode(fi)
	if err != nil {
		msg := fmt.Sprintf("failed to read image: %v", err)
		panic(msg)
	}

	source := gozxing.NewLuminanceSourceFromImage(img)
	bin := gozxing.NewHybridBinarizer(source)
	bbm, err := gozxing.NewBinaryBitmap(bin)

	if err != nil {
		msg := fmt.Sprintf("error during processing: %v", err)
		panic(msg)
	}

	qrReader := qrcode.NewQRCodeMultiReader()
	result, err := qrReader.DecodeMultiple(bbm, nil)
	if err != nil {
		msg := fmt.Sprintf("unable to decode QRCode: %v", err)
		panic(msg)
	}
	strRes := []string{}
	for _, element := range result {
		strRes = append(strRes, element.String())
	}

	res := strings.Join(strRes, "\n")
	fmt.Println(res)
}

func main() {
	CreateQRCode()
	ReadQRCode()
	fmt.Println("Success!")
}
