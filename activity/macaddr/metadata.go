package macaddr

import "github.com/project-flogo/core/data/coerce"


type Output struct {
	Output string `md:"Output"`
	MacAddr string `md:mac_address`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["Output"])
	o.Output = strVal
	strVal1, _ := coerce.ToString(values["MacAddr"])
	o.MacAddr = strVal1
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Output": o.Output,
		"mac_address":o.MacAddr
	}
}
