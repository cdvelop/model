package model

type HttpAdapter interface {
	//action: create,read,update,delete,file
	SendJson(o *Object, in_data []map[string]string, action string, out_resp func([]Response, error))
	SendFormData(object_name string, data map[string]interface{}, out_resp func(Response, error))
}
