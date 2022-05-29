package main

import (
	"fmt"
	"strconv"
	// "reflect"
	// "os"
) 

type State interface {
	apply() (string);
}

type Q1 struct {}
func (Q1) apply() (string) {
	return "q1";
}

type Q2 struct {}
func (Q2) apply() (string) {
	return "q2";
}

type Q3 struct {}
func (Q3) apply() (string) {
	return "q3";
}

var (
	q1 Q1;
	q2 Q2;
	q3 Q3;
	//TODO: Draw transition table here
	trans = map[State][]State{
		q1: {q2, q1},
		q2: {q2, q3},
		q3: {q2, q1},
	}
)

func main() {
	// instring := os.Args[1]
	var input string;
	fmt.Printf("> ")
	fmt.Scanf("%s\n", &input)
	var curr_state State;
	curr_state = q1; // start state

	for _, v := range(input) {
		curr_sym, _ := strconv.Atoi(string(v));
		//fmt.Println(curr_sym, err, reflect.TypeOf(curr_sym))
		new_state := trans[curr_state][curr_sym]
		fmt.Printf("%s --> %s : %d\n",
		curr_state.apply(), 
		new_state.apply(),
		curr_sym)
		curr_state = new_state
	}

	if curr_state == q3 {
		fmt.Println("Congrats, your string ends with 01");
	}else {
		fmt.Println("Nope nope, your string doesn't ends with 01");
	}
}
