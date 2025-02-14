package function

import "reflect"

// IsNil checks if the interface is a nil or a pointer to a nil
func IsNil(i interface{}) bool {
	// See also
	// https://glucn.com/posts/2019-05-20-golang-an-interface-holding-a-nil-value-is-not-nil
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}
