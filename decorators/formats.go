package proto

import (
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
)

func DumpPlain(object interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%+v", object)), nil
}

func DumpJSON(object interface{}) ([]byte, error) {
	return json.Marshal(&object)
}

func DumpYAML(object interface{}) ([]byte, error) {
	return yaml.Marshal(object)
}
