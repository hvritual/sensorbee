package tuple

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestTuple(t *testing.T) {
	var testData = Map{
		"bool":   Bool(true),
		"int":    Int(1),
		"float":  Float(0.1),
		"string": String("homhom"),
		"byte":   Blob([]byte("madmad")),
		"time":   Timestamp(time.Date(2015, time.April, 10, 10, 23, 0, 0, time.UTC)),
		"array":  Array([]Value{String("saysay"), String("mammam")}),
		"map": Map{
			"string": String("homhom"),
		},
	}
	tup := Tuple{
		Data:          testData,
		Timestamp:     time.Date(2015, time.April, 10, 10, 23, 0, 0, time.UTC),
		ProcTimestamp: time.Date(2015, time.April, 10, 10, 24, 0, 0, time.UTC),
		BatchID:       7,
	}

	Convey("Given a Tuple with values in it", t, func() {
		Convey("When deep-copying the Tuple", func() {
			copy := tup.Copy()

			Convey("Then tuple metadata should be the same", func() {
				So(tup.Timestamp, ShouldResemble, copy.Timestamp)
				So(&tup.Timestamp, ShouldNotPointTo, &copy.Timestamp)

				So(tup.ProcTimestamp, ShouldResemble, copy.ProcTimestamp)
				So(&tup.ProcTimestamp, ShouldNotPointTo, &copy.ProcTimestamp)

				So(tup.BatchID, ShouldResemble, copy.BatchID)
				So(&tup.BatchID, ShouldNotPointTo, &copy.BatchID)
			})

			Convey("Then all values should be the same", func() {
				simpleTypes := []string{"bool", "int", "float", "string",
					"array[0]", "map/string"}
				for _, typeName := range simpleTypes {
					a, getErrA := tup.Data.Get(typeName)
					So(getErrA, ShouldBeNil)
					b, getErrB := copy.Data.Get(typeName)
					So(getErrB, ShouldBeNil)
					// objects should have the same value
					So(a, ShouldEqual, b)
					// pointers should not be the same
					So(&a, ShouldNotPointTo, &b)
				}

				complexTypes := []string{"byte", "time"}
				for _, typeName := range complexTypes {
					a, getErrA := tup.Data.Get(typeName)
					So(getErrA, ShouldBeNil)
					b, getErrB := copy.Data.Get(typeName)
					So(getErrB, ShouldBeNil)
					// objects should have the same value
					So(a, ShouldResemble, b)
					// pointers should not be the same
					So(&a, ShouldNotPointTo, &b)
				}
			})
		})
	})
}
