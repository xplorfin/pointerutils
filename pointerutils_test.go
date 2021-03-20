package pointerutils

import (
	"testing"
)

const (
	testNumInt   = 42
	testNumFloat = 42.
	testNumStr   = "42"
)

func TestInts(t *testing.T) {
	t.Run("int", testInt)
	t.Run("int8", testInt8)
	t.Run("int16", testInt16)
	t.Run("int32", testInt32)
	t.Run("int64", testInt64)
}

func testInt(t *testing.T) {
	t.Parallel()

	t.Run("check int pointer is not nil", func(t *testing.T) {
		n := FromInt(testNumInt)
		if n == nil {
			t.Error("int pointer was nil")
		}
	})
	t.Run("check int pointer returns proper value", func(t *testing.T) {
		n := ToInt(FromInt(testNumInt))
		if n != testNumInt {
			t.Errorf("converting int pointer to int returned unexpected value. "+
				"Expected: %[1]d, Actual: %[2]d", testNumInt, n)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		n := ToInt(nil)
		if n != 0 {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func testInt8(t *testing.T) {
	t.Parallel()

	t.Run("check int8 pointer is not nil", func(t *testing.T) {
		n := FromInt8(testNumInt)
		if n == nil {
			t.Error("int8 pointer was nil")
		}
	})
	t.Run("check int8 pointer returns proper value", func(t *testing.T) {
		n := ToInt8(FromInt8(testNumInt))
		if n != testNumInt {
			t.Errorf("converting int8 pointer to int8 returned unexpected value. "+
				"Expected: %[1]d, Actual: %[2]d", testNumInt, n)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		n := ToInt8(nil)
		if n != 0 {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func testInt16(t *testing.T) {
	t.Parallel()

	t.Run("check int16 pointer is not nil", func(t *testing.T) {
		n := FromInt16(testNumInt)
		if n == nil {
			t.Error("int16 pointer was nil")
		}
	})
	t.Run("check int16 pointer returns proper value", func(t *testing.T) {
		n := ToInt16(FromInt16(testNumInt))
		if n != testNumInt {
			t.Errorf("converting int16 pointer to int16 returned unexpected value. "+
				"Expected: %[1]d, Actual: %[2]d", testNumInt, n)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		n := ToInt16(nil)
		if n != 0 {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func testInt32(t *testing.T) {
	t.Parallel()

	t.Run("check int32 pointer is not nil", func(t *testing.T) {
		n := FromInt32(testNumInt)
		if n == nil {
			t.Error("int32 pointer was nil")
		}
	})
	t.Run("check int32 pointer returns proper value", func(t *testing.T) {
		n := ToInt32(FromInt32(testNumInt))
		if n != testNumInt {
			t.Errorf("converting int32 pointer to int32 returned unexpected value. "+
				"Expected: %[1]d, Actual: %[2]d", testNumInt, n)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		n := ToInt32(nil)
		if n != 0 {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func testInt64(t *testing.T) {
	t.Parallel()

	t.Run("check int64 pointer is not nil", func(t *testing.T) {
		n := FromInt64(testNumInt)
		if n == nil {
			t.Error("int64 pointer was nil")
		}
	})
	t.Run("check int64 pointer returns proper value", func(t *testing.T) {
		n := ToInt64(FromInt64(testNumInt))
		if n != testNumInt {
			t.Errorf("converting int64 pointer to int64 returned unexpected value. "+
				"Expected: %[1]d, Actual: %[2]d", testNumInt, n)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		n := ToInt64(nil)
		if n != 0 {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func TestUints(t *testing.T) {
	t.Run("uint", testUint)
	t.Run("uint8", testUint8)
	t.Run("uint16", testUint16)
	t.Run("uint32", testUint32)
	t.Run("uint64", testUint64)
}

func testUint(t *testing.T) {
	t.Parallel()

	t.Run("check uint pointer is not nil", func(t *testing.T) {
		n := FromUint(testNumInt)
		if n == nil {
			t.Error("uint pointer was nil")
		}
	})
	t.Run("check uint pointer returns proper value", func(t *testing.T) {
		n := ToUint(FromUint(testNumInt))
		if n != testNumInt {
			t.Errorf("converting uint pointer to uint returned unexpected value. "+
				"Expected: %[1]d, Actual: %[2]d", testNumInt, n)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		n := ToUint(nil)
		if n != 0 {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func testUint8(t *testing.T) {
	t.Parallel()

	t.Run("check uint8 pointer is not nil", func(t *testing.T) {
		n := FromUint8(testNumInt)
		if n == nil {
			t.Error("uint8 pointer was nil")
		}
	})
	t.Run("check uint8 pointer returns proper value", func(t *testing.T) {
		n := ToUint8(FromUint8(testNumInt))
		if n != testNumInt {
			t.Errorf("converting uint8 pointer to uint8 returned unexpected value. "+
				"Expected: %[1]d, Actual: %[2]d", testNumInt, n)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		n := ToUint8(nil)
		if n != 0 {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func testUint16(t *testing.T) {
	t.Parallel()

	t.Run("check uint16 pointer is not nil", func(t *testing.T) {
		n := FromUint16(testNumInt)
		if n == nil {
			t.Error("uint16 pointer was nil")
		}
	})
	t.Run("check uint16 pointer returns proper value", func(t *testing.T) {
		n := ToUint16(FromUint16(testNumInt))
		if n != testNumInt {
			t.Errorf("converting uint16 pointer to uint16 returned unexpected value. "+
				"Expected: %[1]d, Actual: %[2]d", testNumInt, n)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		n := ToUint16(nil)
		if n != 0 {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func testUint32(t *testing.T) {
	t.Parallel()

	t.Run("check uint32 pointer is not nil", func(t *testing.T) {
		n := FromUint32(testNumInt)
		if n == nil {
			t.Error("uint32 pointer was nil")
		}
	})
	t.Run("check uint32 pointer returns proper value", func(t *testing.T) {
		n := ToUint32(FromUint32(testNumInt))
		if n != testNumInt {
			t.Errorf("converting uint32 pointer to uint32 returned unexpected value. "+
				"Expected: %[1]d, Actual: %[2]d", testNumInt, n)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		n := ToUint32(nil)
		if n != 0 {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func testUint64(t *testing.T) {
	t.Parallel()

	t.Run("check uint64 pointer is not nil", func(t *testing.T) {
		n := FromUint64(testNumInt)
		if n == nil {
			t.Error("uint64 pointer was nil")
		}
	})
	t.Run("check uint64 pointer returns proper value", func(t *testing.T) {
		n := ToUint64(FromUint64(testNumInt))
		if n != testNumInt {
			t.Errorf("converting uint64 pointer to uint64 returned unexpected value. "+
				"Expected: %[1]d, Actual: %[2]d", testNumInt, n)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		n := ToUint64(nil)
		if n != 0 {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func TestFloats(t *testing.T) {
	t.Run("float32", testFloat32)
	t.Run("float64", testFloat64)
}

func testFloat32(t *testing.T) {
	t.Parallel()

	t.Run("check float32 pointer is not nil", func(t *testing.T) {
		f := FromFloat32(testNumFloat)
		if f == nil {
			t.Error("float32 pointer was nil")
		}
	})
	t.Run("check float32 pointer returns proper value", func(t *testing.T) {
		f := ToFloat32(FromFloat32(testNumFloat))
		if f != testNumFloat {
			t.Errorf("converting float32 pointer to float32 returned unexpected value. "+
				"Expected: %[1]f, Actual: %[2]f", testNumFloat, f)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		f := ToFloat32(nil)
		if f != 0. {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func testFloat64(t *testing.T) {
	t.Parallel()

	t.Run("check float64 pointer is not nil", func(t *testing.T) {
		f := FromFloat64(testNumFloat)
		if f == nil {
			t.Error("float64 pointer was nil")
		}
	})
	t.Run("check float32 pointer returns proper value", func(t *testing.T) {
		f := ToFloat64(FromFloat64(testNumFloat))
		if f != testNumFloat {
			t.Errorf("converting float64 pointer to float64 returned unexpected value. "+
				"Expected: %[1]f, Actual: %[2]f", testNumFloat, f)
		}
	})
	t.Run("check nil pointer returns 0", func(t *testing.T) {
		f := ToFloat64(nil)
		if f != 0. {
			t.Error("nil pointer didn't return 0")
		}
	})
}

func TestBool(t *testing.T) {
	t.Parallel()

	t.Run("check bool pointer is not nil", func(t *testing.T) {
		b := FromBool(true)
		if b == nil {
			t.Error("bool pointer was nil")
		}
	})
	t.Run("check bool pointer returns proper value", func(t *testing.T) {
		b := ToBool(FromBool(true))
		if !b {
			t.Errorf("converting bool pointer to bool returned unexpected value. "+
				"Expected: %[1]t, Actual: %[2]t", true, b)
		}
	})
	t.Run("check nil pointer returns false", func(t *testing.T) {
		b := ToBool(nil)
		if b {
			t.Error("nil pointer didn't return false")
		}
	})
}

func TestString(t *testing.T) {
	t.Parallel()

	t.Run("check string pointer is not nil", func(t *testing.T) {
		s := FromString(testNumStr)
		if s == nil {
			t.Error("string pointer was nil")
		}
	})
	t.Run("check string pointer returns proper value", func(t *testing.T) {
		s := ToString(FromString(testNumStr))
		if s != testNumStr {
			t.Errorf("converting string pointer to string returned unexpected value. "+
				"Expected: %[1]s, Actual: %[2]s", testNumStr, s)
		}
	})
	t.Run("check nil pointer returns empty string", func(t *testing.T) {
		s := ToString(nil)
		if s != "" {
			t.Error("nil pointer didn't return empty string")
		}
	})
}
