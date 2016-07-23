package main

import (
	"fmt"
)

type MyFunkyFunc func(int) string

func main() {
	foo := func(int) string {
		return "I have been inovked!!"
	}
	//goExample(&DummyReadCloser{100, foo}, &DummyReadCloser{200, foo})
	goTypeAssertion(&DummyReadCloser{100, foo})
	//goInterface()
}

type ConflictingReadCloser struct {
}

func (drc *ConflictingReadCloser) Read(b []byte) (int, error) {
	return 1, nil
}

func (drc *ConflictingReadCloser) Close() {
	fmt.Println("Closing resource")
}

type ConflictingCloser interface {
	Close()
}

type ReadCloser interface {
	ConflictingCloser
	Read(b []byte) (n int, err error)
}

type DummyReadCloser struct {
	data int
	mff  MyFunkyFunc
}

func (drc *DummyReadCloser) Read(b []byte) (int, error) {
	return 1, nil
}

func (drc *DummyReadCloser) Close() {
	fmt.Printf("Closing resource %d", drc.data)
}

func goExample(r ReadCloser, cc ConflictingCloser) {
	r.Close()
	cc.Close()
	fmt.Printf("%v", r)
	fmt.Printf("%v", cc)
}

func goTypeAssertion(r ReadCloser) {
	if result, ok := r.(*DummyReadCloser); ok {
		fmt.Println(result.mff(10))
		fmt.Printf("Type assertion: %v", result)
	} else {
		fmt.Printf("Wrong type assertions")
	}
}

func goInterface() {
	var t interface{}
	t = callMe()

	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t)
	}
}

func callMe() interface{} {
	return DummyReadCloser{}
}
