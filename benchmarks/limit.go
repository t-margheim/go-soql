package benchmarks

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func getReflectedValueAndType(v interface{}) (reflect.Value, reflect.Type, error) {
	var reflectedValue reflect.Value
	var reflectedType reflect.Type
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Value{}, nil, errors.New("nil value")
		}
		reflectedValue = reflect.Indirect(reflect.ValueOf(v))
	} else {
		reflectedValue = reflect.ValueOf(v)
	}
	reflectedType = reflectedValue.Type()
	return reflectedValue, reflectedType, nil
}
func marshalLimitClause(v interface{}) (string, error) {
	limit, ok := v.(int)
	if !ok {
		return "", errors.New("invalid")
	}
	if limit < 0 {
		return "", errors.New("invalid")
	}

	var limitString string
	if limit > 0 {
		limitString = strconv.Itoa(limit)
	}
	return limitString, nil
}

func marshalLimitClauseAlt(v interface{}) (string, error) {
	reflectedValue, reflectedType, err := getReflectedValueAndType(v)
	if err != nil {
		return "", err
	}

	if reflectedType.Kind() != reflect.Int {
		return "", errors.New("invalid")
	}
	limit := reflectedValue.Int()
	if limit < 0 {
		return "", errors.New("invalid")
	}

	var buff strings.Builder
	if limit > 0 {
		limitString := strconv.FormatInt(limit, 10)
		buff.WriteString(limitString)
	}
	return buff.String(), nil
}
