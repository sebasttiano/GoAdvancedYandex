package main

// JSONData — интерфейс для декодирования JSON.
type JSONData interface {
	DecodeJSON() interface{}
}

// YAMLData — интерфейс для декодирования YAML.
type YAMLData interface {
	DecodeYAML() interface{}
}

type Client struct {
	Data interface{}
}

func (client *Client) Decode(input JSONData) {
	client.Data = input.DecodeJSON()
}

// добавьте тип Adapter и необходимый метод
// ...

type MyAdapter struct {
	ymlData YAMLData
}

func (a *MyAdapter) DecodeJSON() interface{} {
	return a.ymlData.DecodeYAML()
}
func Load(client Client, input YAMLData) {
	adapter := &MyAdapter{
		ymlData: input,
	}
	client.Decode(adapter)
}
