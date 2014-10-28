// Package gunvalue provides some help to work with glib.Value.
package gunvalue

import (
	"github.com/conformal/gotk3/glib"

	"errors"
	"unsafe"
)

//
//----------------------------------------------------------[ ICONS BY ORDER ]--

// Conv converts a glib.Value to a usable go type.
//
type Conv struct {
	data interface{}
	err  error
}

// New converts the glib.Value func result to a usable go value.
// can be used with TreeModel.GetValue
func New(gval *glib.Value, err error) Conv {
	if err != nil {
		return Conv{err: err}
	}
	val, ego := gval.GoValue()
	if ego != nil {
		return Conv{err: ego}
	}
	return Conv{data: val}
}

// Int converts the value to an int.
//
func (c Conv) Int() (int, error) {
	if c.err != nil {
		return 0, c.err
	}
	value, ok := c.data.(int)
	if !ok {
		return 0, errors.New("convert: not int type")
	}
	return value, nil
}

// Pointer converts the value to a pointer.
//
func (c Conv) Pointer() (unsafe.Pointer, error) {
	if c.err != nil {
		return nil, c.err
	}
	value, ok := c.data.(unsafe.Pointer)
	if !ok {
		return nil, errors.New("convert: not unsafe.Pointer type")
	}
	return value, nil
}

// String converts the value to a string.
//
func (c Conv) String() (string, error) {
	if c.err != nil {
		return "", c.err
	}
	value, ok := c.data.(string)
	if !ok {
		return "", errors.New("convert: not string type")
	}
	return value, nil
}