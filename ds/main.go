package main

import (
	passwordgenerator "go_stuff/passwordGenerator"

	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    if true == false {
        arrayStuff();
    }
        mapsStuff();
}

func mapsStuff() {
    var ageMap = map[string]uint8 {
        "Adam": 32,
        "Sarah": 43,
        "Puffie": 25,
    }

    println("Write name of person you want to see the age of:")
    var reader = bufio.NewReader(os.Stdin);
    var input, err = reader.ReadString('\n');

    if err != nil {
        println(err.Error());
    }

    input = strings.TrimSpace(input);

    var age, ok = ageMap[input];

    if ok  {
        fmt.Printf("%v is %v years old \n", input, age);
    } else {
        println("Person not found");
    }

    var i = 0;
    var loopLength, loopErr = passwordgenerator.TakeUserInput("How long do you want the loop");

    if loopErr != nil {
        println(loopErr.Error());
    }

    for i < loopLength {
        i++;
        println("While loop");
    }

    // iterate over map
    for i, v := range ageMap {
        fmt.Printf("%v is %v years old \n", i, v);
    }
}

func arrayStuff () {
    var intArray = []int32{123, 124, 152}

    var length, err = passwordgenerator.TakeUserInput("What is the lenght of the array");

    if err != nil {
        println(err.Error());

        return;
    }

	for i := range length {
	    intArray = append(intArray, int32(i));
	}

    println(len(intArray));

    intArray = append(intArray, intArray...);

    println("Done appending.")
    println(len(intArray));
}
