package qr

import (
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/multi"
	"github.com/makiuchi-d/gozxing/multi/qrcode"
	"github.com/zhukowladimir/qr-interface-for-messenger/utils"
	"image"
	"os"
)

type scanner struct {
	qrReader multi.MultipleBarcodeReader
}

func NewScanner() Scanner {
	return &scanner{
		qrReader: qrcode.NewQRCodeMultiReader(),
	}
}

func (s *scanner) ScanFromImage(imagePath string) ([]*gozxing.Result, error) {
	img, err := s.readImage(imagePath)
	if err != nil {
		return nil, utils.WrapFail("read image", err)
	}

	decoded, err := s.qrReader.DecodeMultipleWithoutHint(img)
	return decoded, utils.WrapFail("decode qr", err)
}

func (*scanner) readImage(imagePath string) (*gozxing.BinaryBitmap, error) {
	imgFile, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer utils.WrapError(imgFile.Close)

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, utils.WrapFail("decode image", err)
	}

	luminanceSource := gozxing.NewLuminanceSourceFromImage(img)
	binarizer := gozxing.NewHybridBinarizer(luminanceSource)
	bitmap, err := gozxing.NewBinaryBitmap(binarizer)
	return bitmap, utils.WrapFail("create bitmap", err)
}
