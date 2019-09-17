package servicenow

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	Username    string `md:"Username,required"`
	Password    string `md:"Password,required"`
	Instanceurl string `md:"Instanceurl,required`
}

type Input struct {
	Content string `md:"content,required"`
}

type Output struct {
	Output string `md:"output"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	//var err error
	i.Content, _ = coerce.ToString(values["content"])
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"content": i.Content,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	o.Output, _ = coerce.ToString(values["output"])

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.Output,
	}
}
