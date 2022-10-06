package mime

type Data struct {
	Type Type
	Data []byte
}

func NewEmptyData(typ Type) *Data {
	return &Data{
		Type: typ,
		Data: nil,
	}
}

func NewData(typ Type, data []byte) *Data {
	return &Data{
		Type: typ,
		Data: data,
	}
}

func NewTextData(text []byte) *Data {
	return &Data{
		Type: TextPlain,
		Data: text,
	}
}

func (m *Data) IsText() bool {
	return m.Type == TextPlain
}

func (m *Data) IsImage() bool {
	return m.Type == ImageBmp || m.Type == ImagePng || m.Type == ImageJpeg
}

func (m *Data) Append(data []byte) {
	if m.Data == nil {
		m.Data = data
	} else {
		m.Data = append(m.Data, data...)
	}
}
