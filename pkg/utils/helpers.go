package utils

import "reflect"

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		// or just v == e
		if reflect.DeepEqual(v, e) {
			return true
		}
	}
	return false
}
