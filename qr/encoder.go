package qr

import (
	rqrcode "github.com/skip2/go-qrcode"
	"github.com/zhukowladimir/qr-interface-for-messenger/utils"
)

type encoder struct {
	recoveryLevel rqrcode.RecoveryLevel
}

func NewEncoder(recoveryLevel rqrcode.RecoveryLevel) Encoder {
	return &encoder{
		recoveryLevel: recoveryLevel,
	}
}

func (e *encoder) EncodeText(text, filePath string) error {
	return utils.WrapFail(
		"encode text to qr",
		rqrcode.WriteFile(text, e.recoveryLevel, 256, filePath),
	)
}
