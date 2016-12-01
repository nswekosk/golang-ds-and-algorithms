package stack

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "gitlab.com/golang-dds-and-algorithms/src/ds/stack"
)

const (
	TEST_NEW_STR   = "Test 1, One string"
	TEST_NEW_STR_2 = "Test 2, SECOND string"
)

func TestPush(t *testing.T) {
	var st Stack
	st = &StringLinkedListStack{}
	st = st.CreateStack()

	Convey("Given no string", t, func() {

		Convey("The size is zero", func() {
			So(st.Size() == 0, ShouldBeTrue)
		})

	})
	st.Push(TEST_NEW_STR)

	Convey("Given one string", t, func() {

		Convey("The size is one", func() {

			So(st.Size() == 1, ShouldBeTrue)
		})

		Convey("Correct peak", func() {

			pk, err := st.Peak()
			So(func() {}, ShouldNotPanicWith, err)
			So(pk == TEST_NEW_STR, ShouldBeTrue)

		})

		Convey("The peak is the same as the pop", func() {
			pk, err := st.Peak()
			p, err := st.Pop()

			So(func() {}, ShouldNotPanicWith, err)
			So(pk == p, ShouldBeTrue)
			So(st.Size() == 0, ShouldBeTrue)
		})
	})

	Convey("Given two string", t, func() {
		st.Push(TEST_NEW_STR)
		st.Push(TEST_NEW_STR_2)

		Convey("The size is two", func() {
			So(st.Size() == 2, ShouldBeTrue)
		})

		Convey("Correct peak", func() {

			pk, err := st.Peak()

			So(func() {}, ShouldNotPanicWith, err)
			So(pk == TEST_NEW_STR, ShouldBeTrue)

		})

		Convey("The peak is the same as the pop", func() {
			p, err := st.Pop()
			So(func() {}, ShouldNotPanicWith, err)
			So(p == TEST_NEW_STR_2, ShouldBeTrue)

			p, err = st.Pop()
			So(func() {}, ShouldNotPanicWith, err)

			So(p == TEST_NEW_STR, ShouldBeTrue)
		})

	})

}
