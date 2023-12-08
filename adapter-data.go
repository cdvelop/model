package model

type SourceData interface {
	SourceData() map[string]string
}

// si nombre del objeto no se ingresa se codifica a json de forma normal
type DataConverter interface {
	EncodeStruct(in any) (result []byte, err string)
	// &out is a pointer
	DecodeStruct(in []byte, out any) (err string)
	//map_in ej []map[string]string or map[string]string
	EncodeMaps(map_in any, object_name ...string) (out []byte, err string)
	DecodeMaps(in []byte, object_name ...string) (out []map[string]string, err string)

	EncodeResponses(requests ...Response) (out []byte, err string)
	DecodeResponses(data []byte) (out []Response, err string)
}
