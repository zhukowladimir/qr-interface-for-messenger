package qr

import "github.com/makiuchi-d/gozxing"

type Scanner interface {
	ScanFromImage(imagePath string) ([]*gozxing.Result, error)
}

type Encoder interface {
	EncodeText(text, filePath string) error
}
