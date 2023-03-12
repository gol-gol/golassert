package golassert

import "testing"

type SampleStruct struct {
	Id int
}

func TestType(t *testing.T) {
	Type(t, "a", "b")
	Type(t, 1, 2)
	Type(t, 1.1, 2.2)
	Type(t, true, false)
	Type(t, []string{"yes"}, []string{"no"})
	Type(t, map[int]string{0: "yes"}, map[int]string{1: "no"})
	Type(t, SampleStruct{Id: 1}, SampleStruct{Id: 2})
}

func TestEqual(t *testing.T) {
	Equal(t, "a", "a")
	Equal(t, 1, 1)
	Equal(t, 1.1, 1.1)
	Equal(t, true, true)
	Equal(t, []string{"yes"}, []string{"yes"})
	Equal(t, map[int]string{0: "yes"}, map[int]string{0: "yes"})
	Equal(t, SampleStruct{Id: 1}, SampleStruct{Id: 1})
}
