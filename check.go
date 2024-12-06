package kettle

import (
	"errors"
	"reflect"
)

func AnyFailed(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func MustTrue(condition bool, message string) error {
	if !condition {
		return errors.New(message)
	}
	return nil
}

func MustFalse(condition bool, message string) error {
	if condition {
		return errors.New(message)
	}
	return nil
}

func MustNil(v interface{}, message string) error {
	if v == nil {
		return nil
	}

	if reflect.ValueOf(v).IsNil() {
		return nil
	}

	return errors.New(message)
}

func MustNotNil(v interface{}, message string) error {
	if v == nil {
		return errors.New(message)
	}

	if reflect.ValueOf(v).IsNil() {
		return errors.New(message)
	}

	return nil
}
