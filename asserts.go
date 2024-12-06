package kettle

import "reflect"

func AssertNil(v interface{}) {
	if v == nil {
		return
	}

	if !reflect.ValueOf(v).IsNil() {
		panic("AssertNil failed")
	}
}

func AssertNotNil(v interface{}) {
	if v == nil {
		panic("AssertNotNil failed")
	}
	if reflect.ValueOf(v).IsNil() {
		panic("AssertNotNil failed")
	}
}

func AssertEqual(v1 interface{}, v2 interface{}) {
	if v1 != v2 {
		panic("AssertEqual failed")
	}
}

func AssertNotEqual(v1 interface{}, v2 interface{}) {
	if v1 == v2 {
		panic("AssertEqual failed")
	}
}

func AssertDeepEqual(v1 interface{}, v2 interface{}) {
	if !reflect.DeepEqual(v1, v2) {
		panic("AssertDeepEqual failed")
	}
}

func AssertDeepNotEqual(v1 interface{}, v2 interface{}) {
	if reflect.DeepEqual(v1, v2) {
		panic("AssertDeepNotEqual failed")
	}
}

func AssertTrue(b bool) {
	if !b {
		panic("AssertTrue failed")
	}
}

func AssertFalse(b bool) {
	if b {
		panic("AssertFalse failed")
	}
}
