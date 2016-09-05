# glisp

## Lisp like functionality in golang

Any values can be used as elements in the lists. Thus you can use your own structs etc. as elements, you don't need to do anything special to include them.

The functions all use interface{} and return interface{} values, this makes it just easier seeing that we can make use of any types.

Currently there is no scanner to convert a string to a list and also no evaluator.

A basic example of what can be done is:

	package main

	import (
		"fmt"
		lsp "github.com/hduplooy/glisp"
	)

	func main() {
		lst1 := lsp.List(1, 2, 3, "Hansel", "Gretel", lsp.List("+", "a", "b"))
		lst2 := lsp.List("diff", lsp.List("cos", lsp.List("^", "x", 2)), "x")
		lst3 := lsp.Append(lst1, lst2)
		lst4 := lsp.Clone(lst1)
		fmt.Println("lst1=" + lsp.ToString(lst1))
		fmt.Println("lst2=" + lsp.ToString(lst2))
		fmt.Println("lst3 Append(lst1,lst2)=" + lsp.ToString(lst3))
		fmt.Println("lst4 Clone(lst1)=" + lsp.ToString(lst4))
		fmt.Println("Nth(lst1,3)=" + lsp.ToString(lsp.Nth(lst1, 3)))
		fmt.Println("NthCdr(lst1,3)=" + lsp.ToString(lsp.NthCdr(lst1, 3)))
		fmt.Println("Head(lst3,4)=" + lsp.ToString(lsp.Head(lst3, 4)))
		fmt.Printf("Length(lst3)=%d\n", lsp.Length(lst3))
		fmt.Println("Tail(lst3,4)=" + lsp.ToString(lsp.Tail(lst3, 4)))
		lst5 := lsp.Map(func(vals []interface{}) interface{} {
			return lsp.List(vals...)
		}, lst1, lsp.Reverse(lst1))
		fmt.Println("lst5 Map=" + lsp.ToString(lst5))
		fmt.Println("Cons(\"a\",2)=" + lsp.ToString(lsp.Cons("a", 2)))
		fmt.Println("MakeList(5,\"*\")=" + lsp.ToString(lsp.MakeList(5, "*")))
		fmt.Println("LastPair(lst1)=" + lsp.ToString(lsp.LastPair(lst1)))
		fmt.Println("Sublist(lst3,5,3)=" + lsp.ToString(lsp.Sublist(lst3, 5, 3)))
		lsp.SetNth(lst1, 2, "XXX")
		fmt.Println("SetNth(lst1,2,\"XXX\")=" + lsp.ToString(lst1))
		lst6 := lsp.Clone(lst3)
		lsp.SetNthCdr(lst6, 4, lsp.List("a", "b", "c"))
		fmt.Println("SetNthCdr(Clone(lst3),4,List(a,b,c))=" + lsp.ToString(lst6))
		fmt.Println("Filter=" + lsp.ToString(lsp.Filter(func(val interface{}) bool {
			return lsp.IsNumber(val)
		}, lst1)))
		fmt.Println("Delete=" + lsp.ToString(lsp.Delete(func(val interface{}) bool {
			return lsp.IsNumber(val)
		}, lst1)))
		fmt.Println("Member=" + lsp.ToString(lsp.Member(func(val interface{}) bool {
			return lsp.IsString(val) && val.(string) == "Hansel"
		}, lst3)))
		fmt.Println("Fold=" + lsp.ToString(lsp.Fold(func(lst1, lst2 interface{}) interface{} {
			if lsp.IsInt(lst1) {
				return lst1.(int) + lst2.(int)
			}
			return lst2
		}, lst3, 0)))
		lst7 := lsp.List(lsp.Cons("john", 16), lsp.Cons("mary", 14), lsp.Cons("bob", 15), lsp.Cons("sarah", 13))
		fmt.Println("lst7=" + lsp.ToString(lst7))
		fmt.Println("Assoc=" + lsp.ToString(lsp.Assoc(func(val1 interface{}) bool {
			return lsp.IsString(val1) && val1.(string) == "mary"
		}, lst7)))
	}
	
This will return the following:

	lst1=(1 2 3 Hansel Gretel (+ a b))
	lst2=(diff (cos (^ x 2)) x)
	lst3 Append(lst1,lst2)=(1 2 3 Hansel Gretel (+ a b) diff (cos (^ x 2)) x)
	lst4 Clone(lst1)=(1 2 3 Hansel Gretel (+ a b))
	Nth(lst1,3)=Hansel
	NthCdr(lst1,3)=(Hansel Gretel (+ a b))
	Head(lst3,4)=(1 2 3 Hansel)
	Length(lst3)=9
	Tail(lst3,4)=((+ a b) diff (cos (^ x 2)) x)
	lst5 Map=((1 (+ a b)) (2 Gretel) (3 Hansel) (Hansel 3) (Gretel 2) ((+ a b) 1))
	Cons("a",2)=(a . 2)
	MakeList(5,"*")=(* * * * *)
	LastPair(lst1)=((+ a b))
	Sublist(lst3,5,3)=((+ a b) diff (cos (^ x 2)))
	SetNth(lst1,2,"XXX")=(1 2 XXX Hansel Gretel (+ a b))
	SetNthCdr(Clone(lst3),4,List(a,b,c))=(1 2 3 Hansel Gretel a b c)
	Filter=(1 2)
	Delete=(XXX Hansel Gretel (+ a b))
	Member=(Hansel Gretel (+ a b) diff (cos (^ x 2)) x)
	Fold=6
	lst7=((john . 16) (mary . 14) (bob . 15) (sarah . 13))
	Assoc=(mary . 14)

## API

## Checking helper functions

### IsNode(val interface{}) bool
IsNode checks if val is a Node

### func IsNil(val interface{}) bool
IsNil is a helper function to check if val is nil

### func IsString(val interface{}) bool
IsString is a helper function to check if val is a string

### IsInt(val interface{}) bool
IsInt is a help function to check if val is an int

### IsInt8(val interface{}) bool
IsInt8 is a help function to check if val is an int8

### IsInt16(val interface{}) bool
IsInt16 is a help function to check if val is an int16

### IsInt32(val interface{}) bool
IsInt32 is a help function to check if val is an int32

### IsInt64(val interface{}) bool
IsInt64 is a help function to check if val is an int64

### IsUint(val interface{}) bool
IsUint is a help function to check if val is an uint

### IsUint8(val interface{}) bool
IsUint8 is a help function to check if val is an uint8

### IsUint16(val interface{}) bool
IsUint16 is a help function to check if val is an uint16

### IsUint32(val interface{}) bool
IsUint32 is a help function to check if val is an uint32

### IsUint64(val interface{}) bool
IsUint64 is a help function to check if val is an uint64

### IsByte(val interface{}) bool
IsByte is a help function to check if val is a byte

### IsFloat64(val interface{}) bool
IsFloat64 is a help function to check if val is a float64

### IsFloat32(val interface{}) bool
IsFloat32 is a help function to check if val is a float32

### IsComplex64(val interface{}) bool
IsComplex64 is a help function to check if val is a complex64

### IsComplex128(val interface{}) bool
IsComplex128 is a help function to check if val is a complex128

### IsNumber(val interface{}) bool
IsNumber is a help function to check if val is any of the numeric types

### IsBool(val interface{}) bool
IsBool is a help function to check if val is a bool

## Creation functions

### MakeList(n int, val interface{}) interface{}
MakeList will generate a list of n elements with each element being val.  
For example MakeList(5,"a") will return ("a" "a" "a" "a" "a")

### Cons(val1, val2 interface{}) interface{}
Cons just creates a new node with Car=val1 and Cdr=val2  
For example:  
Cons(4,nil) will return (4)  
Cons("a",4) will return ("a" . 4)

### List(vals ...interface{}) interface{}
List creates a list from the slice of vals

### Append(lsts ...interface{}) interface{}
Append will concatenate lists together and return a new list
for eq. Append on '(a b c) '(1 2 3) becomes '(a b c 1 2 3)

### CloneNode(lst interface{}) interface{}
CloneNode create a new Node with the same Car and Cdr values

### Clone(lst interface{}) interface{}
Clone will create a new list at the top level for deep cloning look at DeepClone
When a cloned list is changed the original will not be changed

### DeepClone(lst interface{}) interface{}
DeepClone will create a new list, but for each value that is also a list it will be also cloned
When anything in the deep cloned list structure is changed the original is not changed

## Extraction functions

### Car(val interface{}) interface{}
Car returns the Car value of a Node else nil if it is not a Node

### Cdr(val interface{}) interface{}
Cdr returns the Cdr value of a Node else nil if it is not a Node

### Nth(lst interface{}, n int) interface{}
Nth will return the n'th value in the list else nil

### NthCdr(lst interface{}, n int) interface{}
NthCdr will return the n'th applcation of Cdr on the list

### Head(lst interface{}, n int) interface{}
Head returns the first n elements of a list else nil if it is not a node or too short

### Tail(lst interface{}, n int) interface{}
Tail will return the last n elements in the list

### LastPair(lst interface{}) interface{}
LastPair returns the very last Node in the list

### Sublist(lst interface{}, start, items int) interface{}
Sublist returns part of the list from the start position for items number of elements

### Length(lst interface{}) int
Length returns the number of elements in the list

## Conversion functions

### Reverse(lst interface{}) interface{}
Reverse returns a list in reverse order

### ToSlice(lst interface{}) []interface{}
ToSlice will convert a list to a slice

### ToString(lst interface{}) string
ToString will convert the list to a string representation of the list

## Modifier functions

### SetCar(lst interface{}, val interface{})
SetCar will set the Car element if it is a Node

### SetCdr(lst interface{}, val interface{})
SetCdr will set the Cdr element if it is a Node

### SetNth(lst interface{}, n int, val interface{}) interface{}
SetNth will set the Car of n'th value of the list to val

### SetNthCdr(lst interface{}, n int, val interface{}) interface{}
SetNthCdr will set the Cdr of n'th value of the list to val

## Utility functions

### Map(f func([]interface{}) interface{}, lsts ...interface{}) interface{}
Map will apply the function f to elements of lists provided and return a new list
For example:
Map(f2,lst1,lst2,lst3) will apply f2 to the first elements of lst1 to lst3 and then the
  second elements etc. and from the results create a new list

### ForEach(f func([]interface{}), lsts ...interface{})
ForEach will apply the function f to elements of lists provided (nothing is returned)
For example:
ForEach(f2,lst1,lst2,lst3) will apply f2 to the first elements of lst1 to lst3 and then the
  second elements etc.

### Filter(f func(interface{}) bool, lst interface{}) interface{}
Filter will use function f to select which elements to return in a new list

### Delete(f func(interface{}) bool, lst interface{}) interface{}
Delete will use function f to select which elements not to return in a new list

### Member(f func(interface{}) bool, lst interface{}) interface{}
Member will use f to return the first Node for which f returns true on its Car

### Fold(f func(interface{}, interface{}) interface{}, lst1, lst2 interface{}) interface{}
Fold will apply f on each of the elements of lst to lst
   this is the same as f(Nth(n),...f(Car(Car(n)),f(Car(n),lst))) ....)

## Association lists

### Acons(val1, val2, lst interface{}) interface{}
Acons will cons val1 and val2 together and add that to the front of lst for a Association list

### Assoc(f func(interface{}) bool, lst interface{}) interface{}
Assoc will search for the first element where the Car of that element will return true when f is applied

