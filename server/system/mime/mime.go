package mime

import "go.uber.org/zap"

type Type uint8

const (
	None Type = iota
	TextPlain
	ImageBmp
	ImagePng
	ImageJpeg
)

func TypeFromString(s string) Type {
	switch s {
	case "text/plain":
		return TextPlain
	case "image/bmp":
		return ImageBmp
	case "image/png":
		return ImagePng
	case "image/jpeg":
		return ImageJpeg
	}
	return None
}

func (m Type) String() string {
	switch m {
	case TextPlain:
		return "text/plain"
	case ImageBmp:
		return "image/bmp"
	case ImagePng:
		return "image/png"
	case ImageJpeg:
		return "image/jpeg"
	default:
		return "none"
	}
}

func (m Type) GoString() string {
	return m.String()
}

func (m Type) IsText() bool {
	return m == TextPlain
}

func (m Type) IsImage() bool {
	return m == ImageBmp || m == ImagePng || m == ImageJpeg
}

func (m Type) ZapField() zap.Field {
	return zap.String("Type", m.String())
}

func (m Type) ZapFieldPrefix(prefix string) zap.Field {
	return zap.String(prefix+"Type", m.String())
}
