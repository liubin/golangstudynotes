package main

import "fmt"

// Person is an Object to stored as value in the map
type Person struct {
	name string
	age int
}


func main(){
	// declare a map variable :
	// var varName map[KEY_TYPE] VALUE_TYPE
	// for thie example ,key is a string type and value is a Person struct.
	var personMap map[string] Person
	
	// and init the var.
	personMap = make(map[string] Person)

	var p,person Person
	var ok bool	

	p = Person{ "user1" , 20 }

	// put a key-value pair to a map as below:
	personMap[p.name] = p

	p = Person{ "user2" , 25 }
	personMap[p.name] = p

	// get a value will like this:
	// the first returned value is the value if found in the map
	// the second returned value (ok as this example) is bool type
	// which indicating whether the key is exists in the map.
	person, ok = personMap["user1"]
	if ok {
		fmt.Println("Found person " , person.name , " age " , person.age)
	} else {
		fmt.Println("Person not found")
	}

	// also we can remove a key-value pair from a map
	delete(personMap,"user1")

	person, ok = personMap["user1"]
	// ok should be false now.
	if ok {

	} else {
		fmt.Println("Person not found")
	}


	/** the output should be:

	Found person  user1  age  20
	Person not found

	*/


}
