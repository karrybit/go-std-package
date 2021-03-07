package main

import (
	"fmt"
	"reflect"
)

type I interface{}
type S struct {
	i int            `itag: "i_tag_value"`
	a [5]int         `atag: "a_tag_value"`
	s []int          `stag: "s_tag_value`
	m map[int]string `mtag: "m_tag_value"`
	c chan int       `ctag: "c_tag_value`
}

func f(i int) int        { return 0 }
func (s S) f(i int) int  { return 0 }
func (s S) F(i int) int  { return 0 }
func (s *S) g(i int) int { return 0 }
func (s *S) G(i int) int { return 0 }

func main() {
	// Type
	// Value
	// Method
	// Kind
	// StructField
	// StructTag
	_reflect()
}

func _reflect() {
	s := S{i: 10, a: [5]int{1, 2, 3, 4, 5}, s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, m: map[int]string{1: "1", 2: "2", 3: "3"}, c: make(chan int)}
	fmt.Printf("\n=== reflectValue ===\n")
	reflectValue(s)
	fmt.Printf("\n=== reflectStruct ===\n")
	reflectStruct(s)
	fmt.Printf("\n=== reflectMethod ===\n")
	reflectMethod(s)
	fmt.Printf("\n=== reflectInt ===\n")
	reflectInt(s.i)
	fmt.Printf("\n=== reflectArray ===\n")
	reflectArray(s.a)
	fmt.Printf("\n=== reflectSlice ===\n")
	reflectSlice(s.s)
	fmt.Printf("\n=== reflectMap ===\n")
	reflectMap(s.m)
	fmt.Printf("\n=== reflectChan ===\n")
	reflectChan(s.c)
	fmt.Printf("\n=== reflectInterface ===\n")
	reflectInterface()
	fmt.Printf("\n=== reflectFunc ===\n\n")
	reflectFunc()
}
func reflectValue(s S) {
	fmt.Printf("reflectValue:ValueOf(s): %+v\n", reflect.ValueOf(s))
	fmt.Printf("reflectValue:ValueOf(&s): %+v\n", reflect.ValueOf(&s))
	// fmt.Printf("reflectValue:Addr(s): %+v\n", reflect.ValueOf(s).Addr())
	// fmt.Printf("reflectValue:Addr(&s): %+v\n", reflect.ValueOf(&s).Addr())
	// fmt.Printf("reflectValue:Bool(s): %+v\n", reflect.ValueOf(s).Bool())
	// fmt.Printf("reflectValue:Bool(&s): %+v\n", reflect.ValueOf(&s).Bool())
	// fmt.Printf("reflectValue:Bytes(s): %+v\n", reflect.ValueOf(s).Bytes())
	// fmt.Printf("reflectValue:Bytes(&s): %+v\n", reflect.ValueOf(&s).Bytes())
}
func reflectInterface() {
	var i interface{}
	fmt.Printf("reflectInterface:TypeOf: %+v\n", reflect.TypeOf(i))
	fmt.Printf("reflectInterface:TypeOf: %+v\n", reflect.TypeOf(new(interface{})))
	fmt.Printf("reflectInterface:Elem: %+v\n", reflect.TypeOf(new(interface{})).Elem())
}
func reflectStruct(s S) {
	fmt.Printf("TypeOf(s): %+v\n", reflect.TypeOf(s))
	fmt.Printf("TypeOf(&s): %+v\n", reflect.TypeOf(&s))
	fmt.Printf("TypeOf(interface{}(s)): %+v\n", reflect.TypeOf(interface{}(s)))
	fmt.Printf("TypeOf(interface{}(&s)): %+v\n", reflect.TypeOf(interface{}(&s)))
	fmt.Printf("Align(s): %+v\n", reflect.TypeOf(s).Align())
	fmt.Printf("Align(&s): %+v\n", reflect.TypeOf(&s).Align())
	fmt.Printf("FieldAlign(s): %+v\n", reflect.TypeOf(s).FieldAlign())
	fmt.Printf("FieldAlign(&s): %+v\n", reflect.TypeOf(&s).FieldAlign())
	fmt.Println("Method(s)(0)", reflect.TypeOf(s).Method(0)) // NOTE: only public method
	// fmt.Println("Method(s)(1)", reflect.TypeOf(s).Method(1))   // NOTE: only public method, panic because second method is pointer method
	fmt.Println("Method(&s)(0)", reflect.TypeOf(&s).Method(0)) // NOTE: only public method
	fmt.Println("Method(&s)(1)", reflect.TypeOf(&s).Method(1)) // NOTE: only public method
	m, b := reflect.TypeOf(s).MethodByName("f")                // NOTE: only public method, retrun false because f is private
	fmt.Printf("MethodByName(s)(\"f\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(s).MethodByName("F") // NOTE: only public method
	fmt.Printf("MethodByName(s)(\"F\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(s).MethodByName("g") // NOTE: only public method, return false because g is private and pointer method
	fmt.Printf("MethodByName(s)(\"g\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(s).MethodByName("G") // NOTE: only public method, return false because G is pointer method
	fmt.Printf("MethodByName(s)(\"G\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(&s).MethodByName("f") // NOTE: only public method, return false because f is private
	fmt.Printf("MethodByName(&s)(\"f\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(&s).MethodByName("F") // NOTE: only public method
	fmt.Printf("MethodByName(&s)(\"F\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(&s).MethodByName("g") // NOTE: only public method, return false because g is private
	fmt.Printf("MethodByName(&s)(\"g\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(&s).MethodByName("G") // NOTE: only public method
	fmt.Printf("MethodByName(&s)(\"G\"): %+v %+v\n", m, b)
	fmt.Printf("NumMethod(s): %+v\n", reflect.TypeOf(s).NumMethod())
	fmt.Printf("NumMethod(&s): %+v\n", reflect.TypeOf(&s).NumMethod())
	fmt.Printf("Name(s): %+v\n", reflect.TypeOf(s).Name())
	fmt.Printf("Name(&s): %+v\n", reflect.TypeOf(&s).Name())
	fmt.Printf("PkgPath(s): %+v\n", reflect.TypeOf(s).PkgPath())
	fmt.Printf("PkgPath(&s): %+v\n", reflect.TypeOf(&s).PkgPath())
	fmt.Printf("Size(s): %+v\n", reflect.TypeOf(s).Size())
	fmt.Printf("Size(&s): %+v\n", reflect.TypeOf(&s).Size())
	fmt.Printf("String(s): %+v\n", reflect.TypeOf(s).String())
	fmt.Printf("String(&s): %+v\n", reflect.TypeOf(&s).String())
	fmt.Printf("Kind(s): %+v\n", reflect.TypeOf(s).Kind())
	fmt.Printf("Kind(&s): %+v\n", reflect.TypeOf(&s).Kind())
	fmt.Printf("Implements(s): %+v\n", reflect.TypeOf(s).Implements(reflect.TypeOf(new(interface{})).Elem()))
	fmt.Printf("Implements(&s): %+v\n", reflect.TypeOf(&s).Implements(reflect.TypeOf(new(interface{})).Elem()))
	fmt.Printf("AssignableTo(s)(%+v): %+v\n", reflect.TypeOf(0), reflect.TypeOf(s).AssignableTo(reflect.TypeOf(0)))
	fmt.Printf("AssignableTo(&s)(%+v): %+v\n", reflect.TypeOf(0), reflect.TypeOf(&s).AssignableTo(reflect.TypeOf(0)))
	fmt.Printf("Comparable(s): %+v\n", reflect.TypeOf(s).Comparable())
	fmt.Printf("Comparable(&s): %+v\n", reflect.TypeOf(&s).Comparable())
	// fmt.Printf("Bits(s): %+v\n", reflect.TypeOf(s).Bits())               // panic because not Arithmetic type
	// fmt.Printf("Bits(&s): %+v\n", reflect.TypeOf(&s).Bits())             // panic because not Arithmetic type
	// fmt.Printf("ChanDir(s): %+v\n", reflect.TypeOf(s).ChanDir())         // panic because not Chan type
	// fmt.Printf("ChanDir(&s): %+v\n", reflect.TypeOf(&s).ChanDir())       // panic because not Chan type
	// fmt.Printf("IsVariadic(s): %+v\n", reflect.TypeOf(s).IsVariadic())   // panic because not Func type
	// fmt.Printf("IsVariadic(&s): %+v\n", reflect.TypeOf(&s).IsVariadic()) // panic because not Func type
	// fmt.Printf("Elem(s): %+v\n", reflect.TypeOf(s).Elem())               // panic because not Array, Chan, Map, Ptr, Slice
	fmt.Printf("Elem(&s): %+v\n", reflect.TypeOf(&s).Elem())
	fmt.Printf("Field(s)(0): %+v\n", reflect.TypeOf(s).Field(0))
	// fmt.Printf("Field(&s)(0): %+v\n", reflect.TypeOf(&s).Field(0)) // panic because not Struct type
	fmt.Printf("FieldByIndex(s)(%+v): %+v\n", []int{0}, reflect.TypeOf(s).FieldByIndex([]int{0}))
	// fmt.Printf("FieldByIndex(&s)(%+v): %+v\n", []int{0}, reflect.TypeOf(&s).FieldByIndex([]int{0})) // panic because not Struct type
	f, b := reflect.TypeOf(s).FieldByName("i")
	fmt.Printf("FieldByName(s)(\"i\"): %+v %+v\n", f, b)
	// f, b = reflect.TypeOf(&s).FieldByName("i") // panic because not Struct type
	// fmt.Printf("FieldByName(&s)(\"i\"): %+v %+v\n", f, b)
	str := ""
	f, b = reflect.TypeOf(s).FieldByNameFunc(func(s string) bool { str += s; return true })
	fmt.Printf("FieldByNameFunc(s)(func(s string) bool { fmt.Println(s); return true }): %+v %+v %+v\n", str, f, b)
	// str = ""
	// f, b = reflect.TypeOf(&s).FieldByNameFunc(func(s string) bool { str += s; return true }) // panic because not Struct type
	// fmt.Printf("FieldByNameFunc(&s)(func(s string) bool { fmt.Println(s); return true }): %+v %+v %+v\n", str, f, b)
	// fmt.Printf("In(s)(0): %+v\n", reflect.TypeOf(s).In(0))   // panic because not Func type
	// fmt.Printf("In(&s)(0): %+v\n", reflect.TypeOf(&s).In(0)) // panic because not Func type
	// fmt.Printf("Key(s): %+v\n", reflect.TypeOf(s).Key())     // panic because not Map type
	// fmt.Printf("Key(&s): %+v\n", reflect.TypeOf(&s).Key())   // panic because not Map type
	// fmt.Printf("Len(s): %+v\n", reflect.TypeOf(s).Len())     // panic because not Array type
	// fmt.Printf("Len(&s): %+v\n", reflect.TypeOf(&s).Len())   // panic because not Array type
	fmt.Printf("NumField(s): %+v\n", reflect.TypeOf(s).NumField())
	// fmt.Printf("NumField(&s): %+v\n", reflect.TypeOf(&s).NumField()) // panic because not Struct type
	// fmt.Printf("NumIn(s): %+v\n", reflect.TypeOf(s).NumIn())         // panic because not Func type
	// fmt.Printf("NumIn(&s): %+v\n", reflect.TypeOf(&s).NumIn())       // panic because not Func type
	// fmt.Printf("NumOut(s): %+v\n", reflect.TypeOf(s).NumOut())       // panic because not Func type
	// fmt.Printf("NumOut(&s): %+v\n", reflect.TypeOf(&s).NumOut())     // panic because not Func type
	// fmt.Printf("Out(s)(0): %+v\n", reflect.TypeOf(s).Out(0))         // panic because not Func type
	// fmt.Printf("Out(&s)(0): %+v\n", reflect.TypeOf(&s).Out(0))       // panic because not Func type
}
func reflectMethod(s S) {
	F, b := reflect.TypeOf(s).MethodByName("F")
	fmt.Printf("MethodByName(s)(\"F\"): %+v %+v\n", F, b)
	fmt.Printf("%+v\n", F)
	fmt.Printf("TypeOf(F): %+v\n", reflect.TypeOf(F))
	G, b := reflect.TypeOf(&s).MethodByName("G")
	fmt.Printf("MethodByName(s)(\"F\"): %+v %+v\n", G, b)
	fmt.Printf("%+v\n", G)
	fmt.Printf("TypeOf(G): %+v\n", reflect.TypeOf(G))
	fmt.Printf("Align(F): %+v\n", reflect.TypeOf(F).Align())
	fmt.Printf("Align(G): %+v\n", reflect.TypeOf(G).Align())
	fmt.Printf("FieldAlign(F): %+v\n", reflect.TypeOf(F).FieldAlign())
	fmt.Printf("FieldAlign(G): %+v\n", reflect.TypeOf(G).FieldAlign())
	// fmt.Println("Method(F)(0)", reflect.TypeOf(F).Method(0)) // NOTE: only public method, panic because not Struct type
	// fmt.Println("Method(G)(0)", reflect.TypeOf(G).Method(0)) // NOTE: only public method, panic because not Struct type
	m, b := reflect.TypeOf(F).MethodByName("f") // NOTE: only public method, retrun false because f is not method
	fmt.Printf("MethodByName(F)(\"f\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(F).MethodByName("F") // NOTE: only public method, return false because g is not method
	fmt.Printf("MethodByName(F)(\"F\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(G).MethodByName("g") // NOTE: only public method, return false because g is not method
	fmt.Printf("MethodByName(G)(\"g\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(G).MethodByName("G") // NOTE: only public method, return false because G is not method
	fmt.Printf("MethodByName(G)(\"G\"): %+v %+v\n", m, b)
	fmt.Printf("NumMethod(F): %+v\n", reflect.TypeOf(F).NumMethod())
	fmt.Printf("NumMethod(G): %+v\n", reflect.TypeOf(G).NumMethod())
	fmt.Printf("Name(F): %+v\n", reflect.TypeOf(F).Name())
	fmt.Printf("Name(G): %+v\n", reflect.TypeOf(G).Name())
	fmt.Printf("PkgPath(F): %+v\n", reflect.TypeOf(F).PkgPath())
	fmt.Printf("PkgPath(G): %+v\n", reflect.TypeOf(G).PkgPath())
	fmt.Printf("Size(F): %+v\n", reflect.TypeOf(F).Size())
	fmt.Printf("Size(G): %+v\n", reflect.TypeOf(G).Size())
	fmt.Printf("String(F): %+v\n", reflect.TypeOf(F).String())
	fmt.Printf("String(G): %+v\n", reflect.TypeOf(G).String())
	fmt.Printf("Kind(F): %+v\n", reflect.TypeOf(F).Kind())
	fmt.Printf("Kind(G): %+v\n", reflect.TypeOf(G).Kind())
	fmt.Printf("Implements(F): %+v\n", reflect.TypeOf(F).Implements(reflect.TypeOf(new(interface{})).Elem()))
	fmt.Printf("Implements(G): %+v\n", reflect.TypeOf(G).Implements(reflect.TypeOf(new(interface{})).Elem()))
	fmt.Printf("AssignableTo(F)(%+v): %+v\n", reflect.TypeOf(0), reflect.TypeOf(F).AssignableTo(reflect.TypeOf(0)))
	fmt.Printf("AssignableTo(G)(%+v): %+v\n", reflect.TypeOf(0), reflect.TypeOf(G).AssignableTo(reflect.TypeOf(0)))
	fmt.Printf("Comparable(F): %+v\n", reflect.TypeOf(F).Comparable())
	fmt.Printf("Comparable(G): %+v\n", reflect.TypeOf(G).Comparable())
	// fmt.Printf("Bits(F): %+v\n", reflect.TypeOf(F).Bits())                                        // panic because not Arithmetic type
	// fmt.Printf("Bits(G): %+v\n", reflect.TypeOf(G).Bits())                                        // panic because not Arithmetic type
	// fmt.Printf("ChanDir(F): %+v\n", reflect.TypeOf(F).ChanDir())                                  // panic because not Chan type
	// fmt.Printf("ChanDir(G): %+v\n", reflect.TypeOf(G).ChanDir())                                  // panic because not Chan type
	// fmt.Printf("IsVariadic(F): %+v\n", reflect.TypeOf(F).IsVariadic())                            // panic because not Func type
	// fmt.Printf("IsVariadic(G): %+v\n", reflect.TypeOf(G).IsVariadic())                            // panic because not Func type
	// fmt.Printf("Elem(F): %+v\n", reflect.TypeOf(F).Elem())                                        // panic because not Array, Chan, Map, Ptr, Slice
	// fmt.Printf("Elem(G): %+v\n", reflect.TypeOf(G).Elem())                                        // panic because not Array, Chan, Map, Ptr, Slice
	// fmt.Printf("Field(F)(0): %+v\n", reflect.TypeOf(F).Field(0))                                  // panic because not Struct type
	// fmt.Printf("Field(G)(0): %+v\n", reflect.TypeOf(G).Field(0))                                  // panic because not Struct type
	// fmt.Printf("FieldByIndex(F)(%+v): %+v\n", []int{0}, reflect.TypeOf(F).FieldByIndex([]int{0})) // panic because not Struct type
	// fmt.Printf("FieldByIndex(G)(%+v): %+v\n", []int{0}, reflect.TypeOf(G).FieldByIndex([]int{0})) // panic because not Struct type
	// f, b := reflect.TypeOf(F).FieldByName("i")                                                    // panic because not Struct type
	// fmt.Printf("FieldByName(F)(\"i\"): %+v %+v\n", f, b)
	// f, b := reflect.TypeOf(G).FieldByName("i") // panic because not Struct type
	// fmt.Printf("FieldByName(G)(\"i\"): %+v %+v\n", f, b)
	// str := ""
	// f, b = reflect.TypeOf(F).FieldByNameFunc(func(s string) bool { str += s; return true }) // panic because not Struct type
	// fmt.Printf("FieldByNameFunc(F)(func(s string) bool { fmt.Println(s); return true }): %+v %+v %+v\n", str, f, b)
	// str := ""
	// f, b = reflect.TypeOf(G).FieldByNameFunc(func(s string) bool { str += s; return true }) // panic because not Struct type
	// fmt.Printf("FieldByNameFunc(G)(func(s string) bool { fmt.Println(s); return true }): %+v %+v %+v\n", str, f, b)
	// fmt.Printf("In(F)(0): %+v\n", reflect.TypeOf(F).In(0))         // panic because not Func type
	// fmt.Printf("In(G)(0): %+v\n", reflect.TypeOf(G).In(0))         // panic because not Func type
	// fmt.Printf("Key(F): %+v\n", reflect.TypeOf(F).Key())           // panic because not Map type
	// fmt.Printf("Key(G): %+v\n", reflect.TypeOf(G).Key())           // panic because not Map type
	// fmt.Printf("Len(F): %+v\n", reflect.TypeOf(F).Len())           // panic because not Array type
	// fmt.Printf("Len(G): %+v\n", reflect.TypeOf(G).Len())           // panic because not Array type
	// fmt.Printf("NumField(F): %+v\n", reflect.TypeOf(F).NumField()) // panic because not Struct type
	// fmt.Printf("NumField(G): %+v\n", reflect.TypeOf(G).NumField()) // panic because not Struct type
	// fmt.Printf("NumIn(F): %+v\n", reflect.TypeOf(F).NumIn())
	// fmt.Printf("NumIn(G): %+v\n", reflect.TypeOf(G).NumIn())
	// fmt.Printf("NumOut(F): %+v\n", reflect.TypeOf(F).NumOut())
	// fmt.Printf("NumOut(G): %+v\n", reflect.TypeOf(G).NumOut())
	// fmt.Printf("Out(F)(0): %+v\n", reflect.TypeOf(F).Out(0)) // panic because not Func type
	// fmt.Printf("Out(G)(0): %+v\n", reflect.TypeOf(G).Out(0)) // panic because not Func type
}
func reflectInt(i int) {
	fmt.Printf("TypeOf(i): %+v\n", reflect.TypeOf(i))
	fmt.Printf("TypeOf(&i): %+v\n", reflect.TypeOf(&i))
	fmt.Printf("TypeOf(interface{}(i)): %+v\n", reflect.TypeOf(interface{}(i)))
	fmt.Printf("TypeOf(interface{}(&i)): %+v\n", reflect.TypeOf(interface{}(&i)))
	fmt.Printf("Align(i): %+v\n", reflect.TypeOf(i).Align())
	fmt.Printf("Align(&i): %+v\n", reflect.TypeOf(&i).Align())
	fmt.Printf("FieldAlign(i): %+v\n", reflect.TypeOf(i).FieldAlign())
	fmt.Printf("FieldAlign(&i): %+v\n", reflect.TypeOf(&i).FieldAlign())
	// fmt.Println("Method(i)(0)", reflect.TypeOf(i).Method(0))   // NOTE: only public method, panic because not Struct type
	// fmt.Println("Method(i)(1)", reflect.TypeOf(i).Method(1))   // NOTE: only public method, panic because not Struct type
	// fmt.Println("Method(&i)(0)", reflect.TypeOf(&i).Method(0)) // NOTE: only public method, panic because not Struct type
	// fmt.Println("Method(&i)(1)", reflect.TypeOf(&i).Method(1)) // NOTE: only public method, panic because not Struct type
	m, b := reflect.TypeOf(i).MethodByName("f") // NOTE: only public method, return false because not Struct type
	fmt.Printf("MethodByName(i)(\"f\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(i).MethodByName("F") // NOTE: only public method, return false because not Struct type
	fmt.Printf("MethodByName(i)(\"F\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(i).MethodByName("g") // NOTE: only public method, return false because not Struct type
	fmt.Printf("MethodByName(i)(\"g\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(i).MethodByName("G") // NOTE: only public method, return false because not Struct type
	fmt.Printf("MethodByName(i)(\"G\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(&i).MethodByName("f") // NOTE: only public method, return false because not Struct type
	fmt.Printf("MethodByName(&i)(\"f\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(&i).MethodByName("F") // NOTE: only public method, return false because not Struct type
	fmt.Printf("MethodByName(&i)(\"F\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(&i).MethodByName("g") // NOTE: only public method, return false because not Struct type
	fmt.Printf("MethodByName(&i)(\"g\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(&i).MethodByName("G") // NOTE: only public method, return false because not Struct type
	fmt.Printf("MethodByName(&i)(\"G\"): %+v %+v\n", m, b)
	fmt.Printf("NumMethod(i): %+v\n", reflect.TypeOf(i).NumMethod())
	fmt.Printf("NumMethod(&i): %+v\n", reflect.TypeOf(&i).NumMethod())
	fmt.Printf("Name(i): %+v\n", reflect.TypeOf(i).Name())
	fmt.Printf("Name(&i): %+v\n", reflect.TypeOf(&i).Name())
	fmt.Printf("PkgPath(i): %+v\n", reflect.TypeOf(i).PkgPath())
	fmt.Printf("PkgPath(&i): %+v\n", reflect.TypeOf(&i).PkgPath())
	fmt.Printf("Size(i): %+v\n", reflect.TypeOf(i).Size())
	fmt.Printf("Size(&i): %+v\n", reflect.TypeOf(&i).Size())
	fmt.Printf("String(i): %+v\n", reflect.TypeOf(i).String())
	fmt.Printf("String(&i): %+v\n", reflect.TypeOf(&i).String())
	fmt.Printf("Kind(i): %+v\n", reflect.TypeOf(i).Kind())
	fmt.Printf("Kind(&i): %+v\n", reflect.TypeOf(&i).Kind())
	fmt.Printf("Implements(i): %+v\n", reflect.TypeOf(i).Implements(reflect.TypeOf(new(interface{})).Elem()))
	fmt.Printf("Implements(&i): %+v\n", reflect.TypeOf(&i).Implements(reflect.TypeOf(new(interface{})).Elem()))
	fmt.Printf("AssignableTo(i)(%+v): %+v\n", reflect.TypeOf(0), reflect.TypeOf(i).AssignableTo(reflect.TypeOf(0)))
	fmt.Printf("AssignableTo(&i)(%+v): %+v\n", reflect.TypeOf(0), reflect.TypeOf(&i).AssignableTo(reflect.TypeOf(0)))
	fmt.Printf("Comparable(i): %+v\n", reflect.TypeOf(i).Comparable())
	fmt.Printf("Comparable(&i): %+v\n", reflect.TypeOf(&i).Comparable())
	fmt.Printf("Bits(i): %+v\n", reflect.TypeOf(i).Bits())
	// fmt.Printf("Bits(&i): %+v\n", reflect.TypeOf(&i).Bits())             // panic because not Arithmetic type
	// fmt.Printf("ChanDir(i): %+v\n", reflect.TypeOf(i).ChanDir())         // panic because not Chan type
	// fmt.Printf("ChanDir(&i): %+v\n", reflect.TypeOf(&i).ChanDir())       // panic because not Chan type
	// fmt.Printf("IsVariadic(i): %+v\n", reflect.TypeOf(i).IsVariadic())   // panic because not Func type
	// fmt.Printf("IsVariadic(&i): %+v\n", reflect.TypeOf(&i).IsVariadic()) // panic because not Func type
	// fmt.Printf("Elem(i): %+v\n", reflect.TypeOf(i).Elem())               // panic because not Array, Chan, Map, Ptr, Slice
	fmt.Printf("Elem(&i): %+v\n", reflect.TypeOf(&i).Elem())
	// fmt.Printf("Field(i)(0): %+v\n", reflect.TypeOf(i).Field(0))                                    // panic because not Struct type
	// fmt.Printf("Field(&i)(0): %+v\n", reflect.TypeOf(&i).Field(0))                                  // panic because not Struct type
	// fmt.Printf("FieldByIndex(i)(%+v): %+v\n", []int{0}, reflect.TypeOf(i).FieldByIndex([]int{0}))   // panic because not Struct type
	// fmt.Printf("FieldByIndex(&i)(%+v): %+v\n", []int{0}, reflect.TypeOf(&i).FieldByIndex([]int{0})) // panic because not Struct type
	// f, b := reflect.TypeOf(i).FieldByName("i")                                                      // panic because not Struct type
	// fmt.Printf("FieldByName(i)(\"i\"): %+v %+v\n", f, b)
	// f, b = reflect.TypeOf(&i).FieldByName("i") // panic because not Struct type
	// fmt.Printf("FieldByName(&i)(\"i\"): %+v %+v\n", f, b)
	// str := ""
	// f, b = reflect.TypeOf(i).FieldByNameFunc(func(s string) bool { str += s; return true }) // panic because not Struct type
	// fmt.Printf("FieldByNameFunc(s)(func(s string) bool { fmt.Println(s); return true }): %+v %+v %+v\n", str, f, b)
	// str = ""
	// f, b = reflect.TypeOf(&i).FieldByNameFunc(func(s string) bool { str += s; return true }) // panic because not Struct type
	// fmt.Printf("FieldByNameFunc(&i)(func(s string) bool { fmt.Println(s); return true }): %+v %+v %+v\n", str, f, b)
	// fmt.Printf("In(i)(0): %+v\n", reflect.TypeOf(i).In(0))           // panic because not Func type
	// fmt.Printf("In(&i)(0): %+v\n", reflect.TypeOf(&i).In(0))         // panic because not Func type
	// fmt.Printf("Key(i): %+v\n", reflect.TypeOf(i).Key())             // panic because not Map type
	// fmt.Printf("Key(&i): %+v\n", reflect.TypeOf(&i).Key())           // panic because not Map type
	// fmt.Printf("Len(i): %+v\n", reflect.TypeOf(i).Len())             // panic because not Array type
	// fmt.Printf("Len(&i): %+v\n", reflect.TypeOf(&i).Len())           // panic because not Array type
	// fmt.Printf("NumField(i): %+v\n", reflect.TypeOf(i).NumField())   // panic because not Struct type
	// fmt.Printf("NumField(&i): %+v\n", reflect.TypeOf(&i).NumField()) // panic because not Struct type
	// fmt.Printf("NumIn(i): %+v\n", reflect.TypeOf(i).NumIn())         // panic because not Func type
	// fmt.Printf("NumIn(&i): %+v\n", reflect.TypeOf(&i).NumIn())       // panic because not Func type
	// fmt.Printf("NumOut(i): %+v\n", reflect.TypeOf(i).NumOut())       // panic because not Func type
	// fmt.Printf("NumOut(&i): %+v\n", reflect.TypeOf(&i).NumOut())     // panic because not Func type
	// fmt.Printf("Out(i)(0): %+v\n", reflect.TypeOf(i).Out(0))         // panic because not Func type
	// fmt.Printf("Out(&i)(0): %+v\n", reflect.TypeOf(&i).Out(0))       // panic because not Func type
}
func reflectArray(a [5]int)       {}
func reflectSlice(s []int)        {}
func reflectMap(m map[int]string) {}
func reflectChan(c chan int)      {}
func reflectFunc() {
	fmt.Printf("TypeOf(f): %+v\n", reflect.TypeOf(f))
	fmt.Printf("TypeOf(interface{}(f)): %+v\n", reflect.TypeOf(interface{}(f)))
	fmt.Printf("Align(f): %+v\n", reflect.TypeOf(f).Align())
	fmt.Printf("FieldAlign(f): %+v\n", reflect.TypeOf(f).FieldAlign())
	// fmt.Println("Method(f)(0)", reflect.TypeOf(f).Method(0)) // NOTE: only public method, panic because not Struct type
	m, b := reflect.TypeOf(f).MethodByName("f") // NOTE: only public method, retrun false because f is not method
	fmt.Printf("MethodByName(f)(\"f\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(f).MethodByName("F") // NOTE: only public method, return false because g is not method
	fmt.Printf("MethodByName(f)(\"F\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(f).MethodByName("g") // NOTE: only public method, return false because g is not method
	fmt.Printf("MethodByName(f)(\"g\"): %+v %+v\n", m, b)
	m, b = reflect.TypeOf(f).MethodByName("G") // NOTE: only public method, return false because G is not method
	fmt.Printf("MethodByName(f)(\"G\"): %+v %+v\n", m, b)
	fmt.Printf("NumMethod(f): %+v\n", reflect.TypeOf(f).NumMethod())
	fmt.Printf("Name(f): %+v\n", reflect.TypeOf(f).Name())
	fmt.Printf("PkgPath(f): %+v\n", reflect.TypeOf(f).PkgPath())
	fmt.Printf("Size(f): %+v\n", reflect.TypeOf(f).Size())
	fmt.Printf("String(f): %+v\n", reflect.TypeOf(f).String())
	fmt.Printf("Kind(f): %+v\n", reflect.TypeOf(f).Kind())
	fmt.Printf("Implements(f): %+v\n", reflect.TypeOf(f).Implements(reflect.TypeOf(new(interface{})).Elem()))
	fmt.Printf("AssignableTo(f)(%+v): %+v\n", reflect.TypeOf(0), reflect.TypeOf(f).AssignableTo(reflect.TypeOf(0)))
	fmt.Printf("Comparable(f): %+v\n", reflect.TypeOf(f).Comparable())
	// fmt.Printf("Bits(f): %+v\n", reflect.TypeOf(f).Bits())             // panic because not Arithmetic type
	// fmt.Printf("ChanDir(f): %+v\n", reflect.TypeOf(f).ChanDir())       // panic because not Chan type
	fmt.Printf("IsVariadic(f): %+v\n", reflect.TypeOf(f).IsVariadic())
	// fmt.Printf("Elem(f): %+v\n", reflect.TypeOf(f).Elem())                                        // panic because not Array, Chan, Map, Ptr, Slice
	// fmt.Printf("Field(f)(0): %+v\n", reflect.TypeOf(f).Field(0))                                  // panic because not Struct type
	// fmt.Printf("FieldByIndex(f)(%+v): %+v\n", []int{0}, reflect.TypeOf(f).FieldByIndex([]int{0})) // panic because not Struct type
	// f, b := reflect.TypeOf(f).FieldByName("i")                                                    // panic because not Struct type
	// fmt.Printf("FieldByName(f)(\"i\"): %+v %+v\n", f, b)
	// str := ""
	// f, b = reflect.TypeOf(f).FieldByNameFunc(func(s string) bool { str += s; return true }) // panic because not Struct type
	// fmt.Printf("FieldByNameFunc(f)(func(s string) bool { fmt.Println(s); return true }): %+v %+v %+v\n", str, f, b)
	fmt.Printf("In(f)(0): %+v\n", reflect.TypeOf(f).In(0))
	// fmt.Printf("Key(f): %+v\n", reflect.TypeOf(f).Key())           // panic because not Map type
	// fmt.Printf("Len(f): %+v\n", reflect.TypeOf(f).Len())           // panic because not Array type
	// fmt.Printf("NumField(f): %+v\n", reflect.TypeOf(f).NumField()) // panic because not Struct type
	fmt.Printf("NumIn(f): %+v\n", reflect.TypeOf(f).NumIn())
	fmt.Printf("NumOut(f): %+v\n", reflect.TypeOf(f).NumOut())
	fmt.Printf("Out(f)(0): %+v\n", reflect.TypeOf(f).Out(0))
}
