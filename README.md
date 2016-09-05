# glisp

## Lisp like functionality in golang

Any values can be used as elements in the lists. Thus you can use your own structs etc. as elements, you don't need to do anything special to include them.

The functions all use interface{} and return interface{} values, this makes it just easier seeing that we can make use of any types.

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

