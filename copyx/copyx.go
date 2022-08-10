package copyx

import (
	"errors"
	"reflect"
)

func DeepCopy(src, target any) error {
	if src == nil || target == nil {
		return errors.New("err: src and target must be nil")
	}
	tarOfT := reflect.TypeOf(target)
	if tarOfT.Kind() != reflect.Pointer {
		return errors.New("err: target must be pointer")
	}
	srcElemOfT := getTypeElem(reflect.TypeOf(src))
	srcElemOfV := getValueElem(reflect.ValueOf(src))
	tarElemOfT := getTypeElem(tarOfT)
	tarElemOfV := getValueElem(reflect.ValueOf(target))
	switch {
	case srcElemOfT.Kind() == reflect.Struct && tarElemOfT.Kind() == reflect.Struct:
		for i := 0; i < srcElemOfV.NumField(); i++ {
			fieldName := srcElemOfV.Type().Field(i).Name
			tarFieldOfV := tarElemOfV.FieldByName(fieldName)
			srcFieldOfV := srcElemOfV.Field(i)
			if tarFieldOfV.IsValid() && srcFieldOfV.Kind() == tarFieldOfV.Kind() {
				if srcFieldOfV.Kind() == reflect.Struct {
					err := DeepCopy(srcFieldOfV, tarFieldOfV)
					if err != nil {
						return errors.New("err:" + err.Error())
					}
				}
				if tarFieldOfV.CanSet() {
					tarFieldOfV.Set(srcFieldOfV)
				}
			}

		}

	default:
		return errors.New("err: 暂不支持此类型拷贝")
	}
	return nil
}
func getValueElem(src reflect.Value) reflect.Value {
	if src.Kind() == reflect.Pointer {
		return src.Elem()
	}
	return src
}
func getTypeElem(src reflect.Type) reflect.Type {
	if src.Kind() == reflect.Pointer {
		return src.Elem()
	}
	return src
}
