package azuremqtt

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	Broker       string                 `md:"broker,required"` // The broker URL
	Id           string                 `md:"id,required"`     // The id of client
	Username     string                 `md:"username"`        // The user's name
	Password     string                 `md:"password"`        // The user's password
	Store        string                 `md:"store"`           // The store for message persistence
	Topic        string                 `md:"topic,required"`  // The topic to publish to
}

type Input struct {
	Message   string `md:"Message,required"`
	ConnectionString string `md:"ConnectionString,required`
	Type of Operation string `md:"Type of Operation"`	
}

type Output struct {
	Output string `md:"Output"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	//var err error
	i.Message,_ = coerce.ToString(values["Message"])
	i.ConnectionString,_ = coerce.ToString(values["ConnectionString"])
	i.Type_of_Operation,_ = coerce.ToString(values["Type of Operation"])
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Username":    i.Username,
		"Password":    i.Password,
		"Brokerurl": i.Brokerurl,
		"Id": i.Id,
		"Topic": i.Topic,
		"Store": i.Store,
		"Message": i.Message,
		"ConnectionString": i.ConnectionString
		"Type of Operation": i.Type_of_Operation

	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	o.Output, _ = coerce.ToString(values["Output"])

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Output": o.Output,
	}
}
