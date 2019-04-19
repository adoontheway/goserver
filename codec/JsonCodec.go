package codec

type jsonCodec struct {
}

func NewJsonCode() ICodec {
	return &jsonCodec{}
}

func (jc *jsonCodec) Encode(data interface{}) interface{} {
	return nil
}

func (jc *jsonCodec) Decode(data interface{}) interface{} {
	return nil
}

func (jc *jsonCodec) GetType() CodecType {
	return CodecJson
}