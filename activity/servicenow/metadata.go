package servicenow

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	Content     string `md:"Content,required"`
	Username    string `md:"Username,required"`
	Password    string `md:"Password,required"`
	Instanceurl string `md:"Instanceurl,required`
}

type Output struct {
	Output string `md:"Output"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	//var err error
	i.Content, _ = coerce.ToString(values["Content"])
	i.Username, _ = coerce.ToString(values["Username"])
	i.Password, _ = coerce.ToString(values["Password"])
	i.Instanceurl, _ = coerce.ToString(values["Instanceurl"])
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Content":     i.Content,
		"Username":    i.Username,
		"Password":    i.Password,
		"Instanseurl": i.Instanceurl,
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
