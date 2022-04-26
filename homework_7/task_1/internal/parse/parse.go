package parse

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func ParseValuesIn(in interface{}, values map[string]interface{}) error {
	refProd := reflect.Indirect(reflect.ValueOf(in))

	if refProd.Kind() != reflect.Struct {
		return fmt.Errorf("error: in isn't a struct")
	}

	for i := 0; i < refProd.Type().NumField(); i++ {
		fld := refProd.Type().Field(i)
		fldn := strings.ToLower(fld.Name)
		fldv := values[fldn]
		ff := refProd.FieldByName(fld.Name)

		if fldvs, ok := fldv.(string); ok &&
			ff.Type().AssignableTo(reflect.TypeOf(uuid.UUID{})) {

			uu, err := uuid.Parse(fldvs)
			if err != nil {
				err = errors.New("uuid.Parse isn't success")
				return err
			}
			ff.Set(reflect.ValueOf(uu))
		} else if fldvs, ok := fldv.(string); ok &&
			ff.Type().AssignableTo(reflect.TypeOf(int64(0))) {

			numInt, err := strconv.ParseInt(fldvs, 10, 64)
			if err != nil {
				err = errors.New("ParseInt isn't success")
				return err
			}
			ff.Set(reflect.ValueOf(numInt))
		} else {
			ff.Set(reflect.ValueOf(fldv))
		}
	}
	return nil
}
