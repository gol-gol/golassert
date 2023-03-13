package golassert

import "testing"

type SampleStruct struct {
	Id   int
	Name string
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

func TestNotType(t *testing.T) {
	NotType(t, "a", 'a')
	NotType(t, 1, 'a')
	NotType(t, 1.1, 2)
	NotType(t, true, "false")
	NotType(t, []string{"yes"}, []int{1})
	NotType(t, map[int]string{0: "yes"}, map[string]string{"n": "no"})
	NotType(t, SampleStruct{Id: 1}, []byte{})
}

func TestEqual(t *testing.T) {
	Equal(t, "a", "a")
	Equal(t, 1, 1)
	Equal(t, 1.1, 1.1)
	Equal(t, true, true)
	Equal(t, []string{"yes"}, []string{"yes"})
	Equal(t, map[int]string{0: "yes"}, map[int]string{0: "yes"})
	Equal(t, map[int]string{0: "ok", 1: "err", 2: "warn"}, map[int]string{0: "ok", 2: "warn", 1: "err"})
	Equal(t, SampleStruct{Id: 1}, SampleStruct{Id: 1})
	Equal(t, SampleStruct{Id: 1, Name: "Alice"}, SampleStruct{Name: "Alice", Id: 1})
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, "a", "b")
	NotEqual(t, 1, 2)
	NotEqual(t, 1.1, 1.2)
	NotEqual(t, true, false)
	NotEqual(t, []string{"yes"}, []string{"ok"})
	NotEqual(t, map[int]string{0: "yes"}, map[int]string{1: "yes"})
	NotEqual(t, map[int]string{0: "ok", 1: "err", 2: "warn"}, map[int]string{0: "ok", 2: "warn", 3: "err"})
	NotEqual(t, SampleStruct{Id: 1}, SampleStruct{Id: 3})
	NotEqual(t, SampleStruct{Id: 1, Name: "Alice"}, SampleStruct{Name: "Bob", Id: 1})
}
