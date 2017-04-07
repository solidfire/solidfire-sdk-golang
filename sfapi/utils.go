package sfapi

import "github.com/ottemo/mapstructure"

func GetDecoder(result interface{}) (*mapstructure.Decoder, error) {
	return mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		Result:           result,
		WeaklyTypedInput: true})
}
