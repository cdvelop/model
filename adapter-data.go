package model

type SourceData interface {
	SourceData() map[string]string
}

// si nombre del objeto no se ingresa se codifica a json de forma normal
type DataConverter interface {
	//map_in ej []map[string]string or map[string]string
	EncodeMaps(map_in any, object_name ...string) (out []byte, err error)
	DecodeMaps(in []byte, object_name ...string) (out []map[string]string, err error)

	EncodeResponses(requests []Response) ([]byte, error)
	DecodeResponses(data []byte) ([]Response, error)
}
