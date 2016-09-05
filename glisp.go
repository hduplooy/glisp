// github.com/hduplooy/glisp
// Author: Hannes du Plooy
// Revision date: 5 Sep 2016
// Lisp like functionality in golang
// interface{} is used as lists/elements etc. This make it easy to use anything in the lists
// Most of the basic functionality is provided (no scanner or evaluator as yet)
// Equality testing is done by funcs seeing that anything can be used as an element
package glisp

import (
	"fmt"
)

// Node is the way that lists are made up with
// Car is the element for this node
// Cdr is pointing to the next element of the list
//        please note that Cdr can be any other element as well
//        so we can do lists like  (1 2 (3 4 5 . 6) . 7) where the dot indicates that the
//        Cdr is not a list and not nil
type Node struct {
	Car interface{}
	Cdr interface{}
}

// IsNode checks if val is a Node
func IsNode(val interface{}) bool {
	switch val.(type) {
	case *Node:
		return true
	}
	return false
}

// IsNil is a helper function to check if val is nil
func IsNil(val interface{}) bool {
	return val == nil
}

// IsString is a helper function to check if val is a string
func IsString(val interface{}) bool {
	switch val.(type) {
	case string:
		return true
	}
	return false
}

// IsInt is a help function to check if val is an int
func IsInt(val interface{}) bool {
	switch val.(type) {
	case int:
		return true
	}
	return false
}

// IsInt8 is a help function to check if val is an int8
func IsInt8(val interface{}) bool {
	switch val.(type) {
	case int8:
		return true
	}
	return false
}

// IsInt16 is a help function to check if val is an int16
func IsInt16(val interface{}) bool {
	switch val.(type) {
	case int16:
		return true
	}
	return false
}

// IsInt32 is a help function to check if val is an int32
func IsInt32(val interface{}) bool {
	switch val.(type) {
	case int32:
		return true
	}
	return false
}

// IsInt64 is a help function to check if val is an int64
func IsInt64(val interface{}) bool {
	switch val.(type) {
	case int64:
		return true
	}
	return false
}

// IsUint is a help function to check if val is an uint
func IsUint(val interface{}) bool {
	switch val.(type) {
	case uint:
		return true
	}
	return false
}

// IsUint8 is a help function to check if val is an uint8
func IsUint8(val interface{}) bool {
	switch val.(type) {
	case uint8:
		return true
	}
	return false
}

// IsUint16 is a help function to check if val is an uint16
func IsUint16(val interface{}) bool {
	switch val.(type) {
	case uint16:
		return true
	}
	return false
}

// IsUint32 is a help function to check if val is an uint32
func IsUint32(val interface{}) bool {
	switch val.(type) {
	case uint32:
		return true
	}
	return false
}

// IsUint64 is a help function to check if val is an uint64
func IsUint64(val interface{}) bool {
	switch val.(type) {
	case uint64:
		return true
	}
	return false
}

// IsByte is a help function to check if val is a byte
func IsByte(val interface{}) bool {
	switch val.(type) {
	case byte:
		return true
	}
	return false
}

// IsFloat64 is a help function to check if val is a float64
func IsFloat64(val interface{}) bool {
	switch val.(type) {
	case float64:
		return true
	}
	return false
}

// IsFloat32 is a help function to check if val is a float32
func IsFloat32(val interface{}) bool {
	switch val.(type) {
	case float32:
		return true
	}
	return false
}

// IsComplex64 is a help function to check if val is a complex64
func IsComplex64(val interface{}) bool {
	switch val.(type) {
	case complex64:
		return true
	}
	return false
}

// IsComplex128 is a help function to check if val is a complex128
func IsComplex128(val interface{}) bool {
	switch val.(type) {
	case complex128:
		return true
	}
	return false
}

// IsNumber is a help function to check if val is any of the numeric types
func IsNumber(val interface{}) bool {
	return IsByte(val) || IsComplex128(val) || IsComplex64(val) || IsFloat32(val) || IsFloat64(val) || IsInt(val) || IsInt8(val) || IsInt16(val) || IsInt32(val) || IsInt64(val) || IsUint(val) || IsUint8(val) || IsUint16(val) || IsUint32(val) || IsUint64(val)
}

// IsBool is a help function to check if val is a bool
func IsBool(val interface{}) bool {
	switch val.(type) {
	case bool:
		return true
	}
	return false
}

// Cons just creates a new node with Car=val1 and Cdr=val2
func Cons(val1, val2 interface{}) interface{} {
	return &Node{val1, val2}
}

// List creates a list from the slice of vals
func List(vals ...interface{}) interface{} {
	tmp := interface{}(nil)
	for i := len(vals) - 1; i >= 0; i-- {
		tmp = Cons(vals[i], tmp)
	}
	return tmp.(*Node)
}

// Car returns the Car value of a Node else nil if it is not a Node
func Car(val interface{}) interface{} {
	if val == nil {
		return nil
	}
	if IsNode(val) {
		return (val.(*Node)).Car
	}
	return nil
}

// Cdr returns the Cdr value of a Node else nil if it is not a Node
func Cdr(val interface{}) interface{} {
	if val == nil {
		return nil
	}
	if IsNode(val) {
		return (val.(*Node)).Cdr
	}
	return nil
}

// Reverse returns a list in reverse order
func Reverse(lst interface{}) interface{} {
	tmp1 := interface{}(nil)
	for lst != nil && IsNode(lst) {
		tmp1 = Cons(Car(lst), tmp1)
		lst = Cdr(lst)
	}
	return tmp1
}

// CloneNode create a new Node with the same Car and Cdr values
func CloneNode(lst interface{}) interface{} {
	if lst == nil || !IsNode(lst) {
		return lst
	}
	tmp := lst.(*Node)
	return &Node{tmp.Car, tmp.Cdr}
}

// Clone will create a new list at the top level for deep cloning look at DeepClone
// When a cloned list is changed the original will not be changed
func Clone(lst interface{}) interface{} {
	if lst == nil || !IsNode(lst) {
		return lst
	}
	ans := interface{}(nil)
	nxt := interface{}(nil)
	for lst != nil && IsNode(lst) {
		tmp := CloneNode(lst)
		if ans == nil {
			ans = tmp
			nxt = tmp
		} else {
			SetCdr(nxt, tmp)
		}
		nxt = tmp
		lst = Cdr(lst)
	}
	return ans
}

// DeepClone will create a new list, but for each value that is also a list it will be also cloned
// When anything in the deep cloned list structure is changed the original is not changed
func DeepClone(lst interface{}) interface{} {
	if lst == nil || !IsNode(lst) {
		return lst
	}
	ans := interface{}(nil)
	nxt := interface{}(nil)
	for lst != nil && IsNode(lst) {
		tmp := &Node{}
		if IsNode(Car(lst)) {
			tmp.Car = DeepClone(Car(lst))
		} else {
			tmp.Car = Car(lst)
		}
		if ans == nil {
			ans = tmp
			nxt = tmp
		} else {
			SetCdr(nxt, tmp)
		}
		nxt = tmp
		lst = Cdr(lst)
	}
	if lst != nil {
		SetCdr(nxt, Cdr(lst))
	}
	return ans
}

// Append will concatenate lists together and return a new list
// for eq. Append on '(a b c) '(1 2 3) becomes '(a b c 1 2 3)
func Append(lsts ...interface{}) interface{} {
	if len(lsts) == 0 {
		return nil
	}
	if lsts[0] == nil || !IsNode(lsts[0]) {
		return nil
	}
	ans := interface{}(nil)
	nxt := interface{}(nil)
	for _, lst := range lsts {
		for lst != nil && IsNode(lst) {
			tmp := CloneNode(lst)
			if ans == nil {
				ans = tmp
				nxt = tmp
			} else {
				SetCdr(nxt, tmp)
			}
			nxt = tmp
			lst = Cdr(lst)
		}
	}
	return ans
}

// ToString will convert the list to a string representation of the list
func ToString(lst interface{}) string {
	if lst == nil {
		return "nil"
	}
	if !IsNode(lst) {
		return fmt.Sprintf("%v", lst)
	}
	ans := "("
	first := true
	for lst != nil && IsNode(lst) {
		if !first {
			ans += " "
		}
		first = false
		val := Car(lst)
		if val == nil {
			ans += "nil"
		} else if IsNode(val) {
			ans += ToString(val)
		} else {
			ans += fmt.Sprintf("%v", val)
		}
		lst = Cdr(lst)
	}
	if lst != nil {
		ans += " . "
		ans += fmt.Sprintf("%v", lst)
	}
	ans += ")"
	return ans
}

// Nth will return the n'th value in the list else nil
func Nth(lst interface{}, n int) interface{} {
	if lst == nil || !IsNode(lst) {
		return nil
	}
	i := 0
	for ; i < n && lst != nil && IsNode(lst); i++ {
		lst = Cdr(lst)
	}
	if lst == nil || !IsNode(lst) {
		return nil
	}
	return Car(lst)
}

// NthCdr will return the n'th applcation of Cdr on the list
func NthCdr(lst interface{}, n int) interface{} {
	if lst == nil || !IsNode(lst) {
		return nil
	}
	i := 0
	for ; i < n && lst != nil && IsNode(lst); i++ {
		lst = Cdr(lst)
	}
	if lst == nil || !IsNode(lst) {
		return nil
	}
	return lst
}

// SetCar will set the Car element if it is a Node
func SetCar(lst interface{}, val interface{}) {
	if lst == nil || !IsNode(lst) {
		return
	}
	(lst.(*Node)).Car = val
}

// SetCdr will set the Cdr element if it is a Node
func SetCdr(lst interface{}, val interface{}) {
	if lst == nil || !IsNode(lst) {
		return
	}
	(lst.(*Node)).Cdr = val
}

// Head returns the first n elements of a list else nil if it is not a node or too short
func Head(lst interface{}, n int) interface{} {
	if lst == nil || !IsNode(lst) {
		return lst
	}
	ans := interface{}(nil)
	nxt := interface{}(nil)
	for lst != nil && IsNode(lst) && n > 0 {
		tmp := CloneNode(lst)
		if ans == nil {
			ans = tmp
			nxt = tmp
		} else {
			SetCdr(nxt, tmp)
		}
		nxt = tmp
		lst = Cdr(lst)
		n--
	}
	SetCdr(nxt, nil)
	return ans
}

// Map will apply the function f to elements of lists provided and return a new list
// For example:
// Map(f2,lst1,lst2,lst3) will apply f2 to the first elements of lst1 to lst3 and then the
//   second elements etc. and from the results create a new list
func Map(f func([]interface{}) interface{}, lsts ...interface{}) interface{} {
	ans := interface{}(nil)
	nxt := interface{}(nil)
	for {
		val := make([]interface{}, len(lsts))
		done := false
		for i, lst := range lsts {
			if lst == nil || !IsNode(lst) {
				done = true
				break
			} else {
				val[i] = Car(lst)
				lsts[i] = Cdr(lst)
			}
		}
		if done {
			break
		}
		tmp := &Node{f(val), nil}
		if ans == nil {
			ans = tmp
			nxt = tmp
		} else {
			SetCdr(nxt, tmp)
		}
		nxt = tmp
	}
	return ans
}

// ForEach will apply the function f to elements of lists provided (nothing is returned)
// For example:
// ForEach(f2,lst1,lst2,lst3) will apply f2 to the first elements of lst1 to lst3 and then the
//   second elements etc.
func ForEach(f func([]interface{}), lsts ...interface{}) {
	for {
		val := make([]interface{}, len(lsts))
		done := false
		for i, lst := range lsts {
			if lst == nil || !IsNode(lst) {
				done = true
				break
			} else {
				val[i] = Car(lst)
				lsts[i] = Cdr(lst)
			}
		}
		if done {
			break
		}
		f(val)
	}
}

// Filter will use function f to select which elements to return in a new list
func Filter(f func(interface{}) bool, lst interface{}) interface{} {
	if lst == nil || !IsNode(lst) {
		return nil
	}
	ans := interface{}(nil)
	nxt := interface{}(nil)
	for lst != nil && IsNode(lst) {
		if f(Car(lst)) {
			tmp := &Node{Car(lst), nil}
			if ans == nil {
				ans = tmp
				nxt = tmp
			} else {
				SetCdr(nxt, tmp)
			}
			nxt = tmp
		}
		lst = Cdr(lst)
	}
	return ans
}

// Delete will use function f to select which elements not to return in a new list
func Delete(f func(interface{}) bool, lst interface{}) interface{} {
	if lst == nil || !IsNode(lst) {
		return nil
	}
	ans := interface{}(nil)
	nxt := interface{}(nil)
	for lst != nil && IsNode(lst) {
		if !f(Car(lst)) {
			tmp := &Node{Car(lst), nil}
			if ans == nil {
				ans = tmp
				nxt = tmp
			} else {
				SetCdr(nxt, tmp)
			}
			nxt = tmp
		}
		lst = Cdr(lst)
	}
	return ans
}

// Length returns the number of elements in the list
func Length(lst interface{}) int {
	if lst == nil || !IsNode(lst) {
		return 0
	}
	cnt := 0
	for lst != nil && IsNode(lst) {
		cnt++
		lst = Cdr(lst)
	}
	return cnt
}

// Tail will return the last n elements in the list
func Tail(lst interface{}, n int) interface{} {
	if lst == nil || !IsNode(lst) {
		return nil
	}
	l := Length(lst)
	return NthCdr(lst, l-n)
}

// MakeList will generate a list of n elements with each element being val
func MakeList(n int, val interface{}) interface{} {
	if n <= 0 {
		return nil
	}
	ans := interface{}(nil)
	nxt := interface{}(nil)
	for ; n > 0; n-- {
		tmp := &Node{val, nil}
		if ans == nil {
			ans = tmp
			nxt = tmp
		} else {
			SetCdr(nxt, tmp)
		}
		nxt = tmp
	}
	return ans
}

// LastPair returns the very last Node in the list
func LastPair(lst interface{}) interface{} {
	if lst == nil || !IsNode(lst) {
		return nil
	}
	for Cdr(lst) != nil && IsNode(Cdr(lst)) {
		lst = Cdr(lst)
	}
	return lst
}

// Sublist returns part of the list from the start position for items number of elements
func Sublist(lst interface{}, start, items int) interface{} {
	return Head(NthCdr(lst, start), items)
}

// SetNth will set the Car of n'th value of the list to val
func SetNth(lst interface{}, n int, val interface{}) interface{} {
	tmp := NthCdr(lst, n)
	SetCar(tmp, val)
	return lst
}

// SetNthCdr will set the Cdr of n'th value of the list to val
func SetNthCdr(lst interface{}, n int, val interface{}) interface{} {
	tmp := NthCdr(lst, n)
	SetCdr(tmp, val)
	return lst
}

// Member will use f to return the first Node for which f returns true on its Car
func Member(f func(interface{}) bool, lst interface{}) interface{} {
	if f == nil || lst == nil {
		return nil
	}
	for lst != nil && IsNode(lst) {
		if f(Car(lst)) {
			return lst
		}
		lst = Cdr(lst)
	}
	return nil
}

// Acons will cons val1 and val2 together and add that to the front of lst for a Association list
func Acons(val1, val2, lst interface{}) interface{} {
	return Cons(Cons(val1, val2), lst)
}

// Assoc will search for the first element where the Car of that element will return true when f is applied
func Assoc(f func(interface{}) bool, lst interface{}) interface{} {
	if lst == nil || !IsNode(lst) {
		return nil
	}
	for lst != nil && IsNode(lst) {
		if IsNode(Car(lst)) && f(Car(Car(lst))) {
			return Car(lst)
		}
		lst = Cdr(lst)
	}
	return nil
}

// ToSlice will convert a list to a slice
func ToSlice(lst interface{}) []interface{} {
	ans := make([]interface{}, Length(lst))
	for i := 0; lst != nil && IsNode(lst); i, lst = i+1, Cdr(lst) {
		ans[i] = Car(lst)
	}
	return ans
}

// Fold will apply f on each of the elements of lst to lst
//    this is the same as f(Nth(n),...f(Car(Car(n)),f(Car(n),lst))) ....)
func Fold(f func(interface{}, interface{}) interface{}, lst1, lst2 interface{}) interface{} {
	if f == nil || lst1 == nil {
		return nil
	}
	for lst1 != nil && IsNode(lst1) {
		lst2 = f(Car(lst1), lst2)
		lst1 = Cdr(lst1)
	}
	return lst2
}
