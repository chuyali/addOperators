package cmd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddOperators(t *testing.T) {
	Convey("testCase1", t, func() {
		a := "123"
		b := int64(6)
		So(AddOperators(a, b), ShouldResemble, []string{"1+2+3", "1*2*3"})
	})
	Convey("testCase2", t, func() {
		a := "232"
		b := int64(8)
		So(AddOperators(a, b), ShouldResemble, []string{"2+3*2", "2*3+2"})
	})
	Convey("testCase3", t, func() {
		a := "105"
		b := int64(5)
		So(AddOperators(a, b), ShouldResemble, []string{"1*0+5", "10-5"})
	})
	Convey("testCase4", t, func() {
		a := "00"
		b := int64(0)
		So(AddOperators(a, b), ShouldResemble, []string{"0+0", "0-0", "0*0"})
	})
}
