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
For example glisp.MakeList(5,"a") will return ("a" "a" "a" "a" "a")

### Cons(val1, val2 interface{}) interface{}
Cons just creates a new node with Car=val1 and Cdr=val2  
For example:  
* glisp.Cons(4,nil) will return (4)
* glisp.Cons("a",4) will return ("a" . 4)

### List(vals ...interface{}) interface{}
List creates a list from the slice of vals  
For example:
* glisp.List(1,2,3,4) will return (1 2 3 4)
* glisp.List("+",1,glisp.List("-","a",10)) will return ("+" 1 ("-" "a" 10))

### Append(lsts ...interface{}) interface{}
Append will concatenate lists together and return a new list.  
For example:
* glisp.Append(glisp.List("a","b","c"),glisp.List(1,2,3)) returns (("a" "b" "c") (1 2 3))

### CloneNode(lst interface{}) interface{}
CloneNode create a new Node with the same Car and Cdr values

### Clone(lst interface{}) interface{}
Clone will create a new list at the top level for deep cloning look at DeepClone.  
When a cloned list is changed the original will not be changed

### DeepClone(lst interface{}) interface{}
DeepClone will create a new list, but for each value that is also a list it will be also cloned.  
When anything in the deep cloned list structure is changed the original is not changed

## Extraction functions

### Car(val interface{}) interface{}
Car returns the Car value of a Node else nil if it is not a Node.  
For example:

    lst1 := glisp.List("a","b","c")  
    lst2 := glisp.Car(lst1) // Returns "a"

### Cdr(val interface{}) interface{}
Cdr returns the Cdr value of a Node else nil if it is not a Node.  
For example:

    lst1 := glisp.List("a","b","c")  
    lst2 := glisp.Cdr(lst1) // Returns ("b" "c")

### Nth(lst interface{}, n int) interface{}
Nth will return the n'th value in the list else nil.  
For example:

    glisp.Nth(glisp.List("a","b","c","d"),2) // will return "c"

### NthCdr(lst interface{}, n int) interface{}
NthCdr will return the n'th applcation of Cdr on the list.  
For example:  

    glisp.NthCdr(glisp.List("a","b","c","d"),2) // will return ("d")

### Head(lst interface{}, n int) interface{}
Head returns the first n elements of a list else nil if it is not a node or too short.  
For example:

    glisp.Head(glisp.List("a","b","c","d"),2) // will return ("a" "b")

### Tail(lst interface{}, n int) interface{}
Tail will return the last n elements in the list.  
For example:  

    glisp.Tail(glisp.List("a","b","c","d"),2) // will return ("c" "d")

### LastPair(lst interface{}) interface{}
LastPair returns the very last Node in the list.  
For example:
* glisp.LastPair(glisp.List("a","b","c","d")) will return ("d")
* glisp.LastPair(glisp.List("a","b","c",glisp.Cons("d","e"))) will return ("d". "e")

### Sublist(lst interface{}, start, items int) interface{}
Sublist returns part of the list from the start position for items number of elements.  
For example:  

    glisp.Sublist(glisp.List("a","b","c","d"),1,2) // will return ("b" "c")

### Length(lst interface{}) int
Length returns the number of elements in the list.  
For example:  
glisp.Length(glisp.List("a","b","c","d")) will return 4

## Conversion functions

### Reverse(lst interface{}) interface{}
Reverse returns a list in reverse order.  
For example:  

    glisp.Reverse(glisp.List("a","b","c","d")) // will return ("d" "c" "b" "a")

### ToSlice(lst interface{}) []interface{}
ToSlice will convert a list to a slice.  
For example:  

    glisp.ToSlice(glisp.List("a","b","c","d")) //will return []interface{}{"a","b","c","d"}

### ToString(lst interface{}) string
ToString will convert the list to a string representation of the list.  
For example:  
glisp.ToString(glisp.List("+",1,glisp.List("-","a",10))) will return "(+ 1 (- a 10))"

## Modifier functions

### SetCar(lst interface{}, val interface{})
SetCar will set the Car element if it is a Node.  
For example:

    lst1 := glisp.Cons("a","b")  
    glisp.SetCar(lst1,"c")  
    fmt.Println(glisp.ToString(lst1)) // Will produce "(c b)"

### SetCdr(lst interface{}, val interface{})
SetCdr will set the Cdr element if it is a Node
For example:

    lst1 := glisp.Cons("a","b")  
    glisp.SetCdr(lst1,"c")  
    fmt.Println(glisp.ToString(lst1)) // Will produce "(a c)"

### SetNth(lst interface{}, n int, val interface{}) interface{}
SetNth will set the Car of n'th value of the list to val
For example:

    lst1 := glisp.List("a","b","c","d","e")  
    glisp.SetNth(lst1,2,"z")  
    fmt.Println(glisp.ToString(lst1)) // Will produce "(a b z d e)"

### SetNthCdr(lst interface{}, n int, val interface{}) interface{}
SetNthCdr will set the Cdr of n'th value of the list to val
For example:

    lst1 := glisp.List("a","b","c","d","e")  
    glisp.SetNthCdr(lst1,2,glisp.List("x","y","z"))  
    fmt.Println(glisp.ToString(lst1)) // Will produce "(a b c x y z)"

## Utility functions

### Map(f func([]interface{}) interface{}, lsts ...interface{}) interface{}
Map will apply the function f to elements of lists provided and return a new list.  
For example:

    f := func{vals []interface{}) interface{} {  
      if len(vals)>=2 {  
          return glisp.Cons(vals[0],vals[1])  
      }  
      return nil  
    }  
    lst1 := glisp.Map(f,glisp.List("a","b","c"),glisp.List(1,2,3)) // This returns (("a" . 1) ("b" . 2) ("c" . 3))

### ForEach(f func([]interface{}), lsts ...interface{})
ForEach will apply the function f to elements of lists provided (nothing is returned).  
For example:

    f := func{vals []interface{}) {  
      if len(vals)>=2 {
          fmt.Printf("%v - %v\n",vals[0],vals[1])
      }  
    }  
    glisp.ForEach(f,glisp.List("a","b","c"),glisp.List(1,2,3)) 
    // This will output:
    // a - 1
    // b - 2
    // c - 3

### Filter(f func(interface{}) bool, lst interface{}) interface{}
Filter will use function f to select which elements to return in a new list. 
For example:

    f := func{val interface{}) bool {  
        return glisp.IsNumber(val)
    }  
    lst1 := glisp.Filter(f,glisp.List("a",5,"b","c",1,"z",2,3)) // This returns (5 1 2 3)

### Delete(f func(interface{}) bool, lst interface{}) interface{}
Delete will use function f to select which elements not to return in a new list. 
For example:

    f := func{val interface{}) bool {  
        return glisp.IsNumber(val)
    }  
    lst1 := glisp.Delete(f,glisp.List("a",5,"b","c",1,"z",2,3)) // This returns ("a" "b" "c" "z")

### Member(f func(interface{}) bool, lst interface{}) interface{}
Member will use f to return the first Node for which f returns true on its Car.  
For example:

    f := func(val interface{}) bool {
        return glisp.IsString(val) && val.(string) == "c"
    }
    lst1 := glisp.List("a","b","c","d","e")
    fmt.Println(glisp.ToString(glisp.Member(f,lst1))) // This prints "(c d e)"

### Fold(f func(interface{}, interface{}) interface{}, lst1, lst2 interface{}) interface{}
Fold will apply f to the first element of lst1 and all of lst2 and the result will then be used to apply f to the second value of lst1 and the returned value, etc.  
This is the same as f(Nth(n),...f(Car(Car(n)),f(Car(n),lst))) ....)  
For example:  

    f := func(val1, val2 interface{}) interface{} {
        if glisp.IsInt(val1) && glisp.IsInt(val2) {
            return val1.(int) + val2.(int) // Just return the sum of the two arguments
        }
        return 0
    }
    lst1 := glisp.List(1,2,3,4,5)
    fmt.Println(ToString(glisp.Fold(f,lst1,0))) // This will output 15
    // Fold will first call f(1,0)  where 0 was provided by the call to fold
    //  this produces 1 because 1+0 = 1
    // Now Fold calls f(2,1) where 1 was returned in the previous step
    //  this produces 3 because 2+1 = 3
    // Now Fold calls f(3,3) which produces 6
    // Then Fold calls f(4,6) which produces 10
    // And lastly Fold calls f(5,10) which returns 15

## Association lists

### Acons(val1, val2, lst interface{}) interface{}
Acons will cons val1 and val2 together and add that to the front of lst for a Association list.  
For example:  

    // Let's assume lst1 holds the list (("a" . 1) ("b" . 2))
    lst1 = glisp.Acons("c",3,lst1) // Now lst1 holds (("c" . 3) ("a" . 1) ("b" . 2))

### Assoc(f func(interface{}) bool, lst interface{}) interface{}
Assoc will search for the first element where the Car of that element will return true when f is applied.
For example:

    // Let's assum lst1 holds the list (("a" . 1) ("b" . 2) ("c" . 3) ("d" . 4))
    f := func(val interface{}) bool {
        return IsString(val) && val.(string)=="b"
    }
    fmt.Println(glisp.Assoc(f,lst1)) // This will print "(b 2)"
    
