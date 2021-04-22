package zero

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
)

func WriteToFile(v interface{}) error {
	t := reflect.TypeOf(v)
	i := Zero(t)
	bytes, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(t.Name()+".json", bytes, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
