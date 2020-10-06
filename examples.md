package main

import "fmt"

const (
	first  = 1
	second = "second"
	third  = iota
	four   = iota
	five   = iota + 6
	six    = 2 << iota
	seven
)

const (
	eight = iota
	nine
)

const pi = 3.1415

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	var firstNameP *string = new(string)
	*firstNameP = "Arthur"
	fmt.Println(*firstNameP)

	lastName := "Corzo"
	fmt.Println(lastName)

	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "Tricia"
	fmt.Println(ptr, *ptr)

	const pi = 3.14
	fmt.Println(pi)

	const con = 3
	fmt.Println(c + 3)
	fmt.Println(c + 1.3)

	fmt.Println(pi)
	fmt.Println(first, second)
	fmt.Println(third, four)
	fmt.Println(five, six)
	fmt.Println(seven)

	fmt.Println(eight, nine)

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])

	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(arr1[1])

	slice := arr[:]
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	slice1 = append(slice1, 4, 42, 27)
	fmt.Println(slice1)

	s2 := slice1[1:]
	s3 := slice1[:2]
	s4 := slice1[1:2]
	fmt.Println(s2, s3, s4)

	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)
	fmt.Println(u.FirstName)

	u2 := user{ID: 1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	fmt.Println(u2)

	wellknownPorts := map[string]int{"http": 80, "https": 443}
	for k := range wellKnownPorts {
		println(k)
	}

	panic("something bad just happened")

}
