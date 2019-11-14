package cosmodb

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	database     string `md:"Database,required"`
	username    string `md:"Username,required"`
	password    string `md:"Password,required"`
	url 		string `md:"Url,required"`
}

type Output struct {
	Output string `md:"Output"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	//var err error
	i.database, _ = coerce.ToString(values["database"])
	i.username, _ = coerce.ToString(values["username"])
	i.password, _ = coerce.ToString(values["password"])
	i.url, _ = coerce.ToString(values["url"])

	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"database":     i.Database,
		"Username":    i.Username,
		"Password":    i.Password,
		"Instanseurl": i.Url,
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
