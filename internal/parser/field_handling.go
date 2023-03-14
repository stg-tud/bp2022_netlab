package parser

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	logger "github.com/gookit/slog"
)

func parseInterfaceInput(fieldValue reflect.Value, targetFieldType reflect.StructField, targetFieldValue reflect.Value) error {
	switch targetFieldType.Type.Kind() {
	case reflect.String:
		stringVal, fits := fieldValue.Interface().(string)
		if !fits {
			return fmt.Errorf("invalid input data type -- expected %s", targetFieldType.Type.Name())
		}
		targetFieldValue.SetString(stringVal)

	case reflect.Int, reflect.Int32, reflect.Int64:
		intVal, fits := fieldValue.Interface().(int64)
		if !fits {
			return fmt.Errorf("invalid input data type -- expected %s", targetFieldType.Type.Name())
		}
		targetFieldValue.SetInt(intVal)

	case reflect.Uint, reflect.Uint32, reflect.Uint64:
		intVal, fits := fieldValue.Interface().(int64)
		if !fits || intVal < 0 {
			return fmt.Errorf("invalid input data type -- expected %s", targetFieldType.Type.Name())
		}
		targetFieldValue.SetUint(uint64(intVal))

	case reflect.Float32, reflect.Float64:
		floatVal, fits := fieldValue.Interface().(float64)
		if !fits {
			return fmt.Errorf("invalid input data type -- expected %s", targetFieldType.Type.Name())
		}
		targetFieldValue.SetFloat(floatVal)

	case reflect.Bool:
		boolVal, fits := fieldValue.Interface().(bool)
		if fits {
			targetFieldValue.SetBool(boolVal)
			return nil
		}
		intVal, fits := fieldValue.Interface().(int64)
		if !fits {
			return fmt.Errorf("invalid input data type -- expected %s", targetFieldType.Type.Name())
		}
		boolVal, err := strconv.ParseBool(strconv.FormatInt(int64(intVal), 10))
		if err != nil {
			return fmt.Errorf("invalid input data type -- expected %s", targetFieldType.Type.Name())
		}
		targetFieldValue.SetBool(boolVal)
	}
	return nil
}

func parseDefaultValue(fieldType reflect.StructField, fieldValue reflect.Value, targetFieldType reflect.StructField, targetFieldValue reflect.Value) error {
	defaultValue, hasDefault := fieldType.Tag.Lookup("default")
	if !hasDefault {
		logger.Trace("Field has no default value.")
		isRequiredString, hasInfo := fieldType.Tag.Lookup("required")
		if !hasInfo {
			logger.Tracef("Field is neither required nor has it a default value. Just skipping it.")
			return nil
		}
		isRequired, err := strconv.ParseBool(isRequiredString)
		if err != nil {
			return err
		}
		if isRequired {
			return errors.New("field is required")
		}
		return nil
	}

	switch targetFieldType.Type.Kind() {
	case reflect.String:
		targetFieldValue.SetString(defaultValue)

	case reflect.Int, reflect.Int32, reflect.Int64:
		intVal, err := strconv.ParseInt(defaultValue, 10, 64)
		if err != nil {
			return err
		}
		targetFieldValue.SetInt(intVal)

	case reflect.Uint, reflect.Uint32, reflect.Uint64:
		uintVal, err := strconv.ParseUint(defaultValue, 10, 64)
		if err != nil {
			return err
		}
		targetFieldValue.SetUint(uintVal)

	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(defaultValue, 32)
		if err != nil {
			return err
		}
		targetFieldValue.SetFloat(floatVal)

	case reflect.Bool:
		boolVal, err := strconv.ParseBool(defaultValue)
		if err != nil {
			return err
		}
		targetFieldValue.SetBool(boolVal)
	}
	return nil
}

func handleField(fieldType reflect.StructField, fieldValue reflect.Value, targetFieldType reflect.StructField, targetFieldValue reflect.Value) error {
	// input and output have same type, just copy value
	if fieldType.Type == targetFieldType.Type {
		targetFieldValue.Set(fieldValue)
		return nil
	}

	// input has interface/any type. Handle required and defaults
	if fieldType.Type.Kind() == reflect.Interface {
		var err error
		if fieldValue.IsZero() || !fieldValue.IsValid() {
			logger.Debugf("Trying to use default value for field \"%s\".", fieldType.Name)
			err = parseDefaultValue(fieldType, fieldValue, targetFieldType, targetFieldValue)
		} else {
			err = parseInterfaceInput(fieldValue, targetFieldType, targetFieldValue)
		}
		if err != nil {
			return fmt.Errorf("error parsing field \"%s\": %s", fieldType.Name, err)
		}
	}
	return nil
}

func fillDefaults[IN any, OUT any](input IN) (OUT, error) {
	var output OUT

	tpe := reflect.TypeOf(input)
	val := reflect.ValueOf(input)
	for i := 0; i < tpe.NumField(); i++ {
		fieldType := tpe.Field(i)
		fieldValue := val.Field(i)
		targetFieldType, found := reflect.TypeOf(output).FieldByName(fieldType.Name)
		if !found {
			logger.Tracef("Field \"%s\" given but not expected", fieldType.Name)
			continue
		}
		targetFieldValue := reflect.ValueOf(&output).Elem().FieldByName(fieldType.Name)
		err := handleField(fieldType, fieldValue, targetFieldType, targetFieldValue)
		if err != nil {
			return output, err
		}
	}

	return output, nil
}
