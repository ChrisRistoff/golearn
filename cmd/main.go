package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
    println(generateRandomNumberInRange(1, 10), "Random number");

    var result, remainder, err = divide(10, 0);

    if (err != nil) {
        println(err.Error());
    } else {
        fmt.Printf("Remainder is %v, result is %v ", remainder, result);
    }

}

func divide(num1 int, num2 int) (int, int, error) {
    var err error;

    if num1 == 0 || num2 == 0 {
        err = errors.New("Cannot divide by 0");
        return 0, 0, err;
    }

    var result = num1 / num2;
    var remainder = num1 % num2;

    return result, remainder, err;
}

func vars(){
    var intNum uint8 = 255;
    intNum++;
    fmt.Println(intNum, " - uint8 overflow");

    var floatNum32 float32 = 12333.12333;
    var floatNum64 float64 = 12333.12333;
    fmt.Println(floatNum32, " - float32");
    fmt.Println(floatNum64, " - float64");

    var myString string = "asdasdas";
    fmt.Println(len(myString));

    var myRune rune = 'a';
    fmt.Println(myRune); // 97 - unicode

    println(string(97)); // turn unicode to a char(string), can use a rune

    var myBoolean bool = true;
    fmt.Println(myBoolean);

    var var1, var2 = 5, 6;
    println(var1, var2);

    var string1 string = "string one";
    var string2 string = "string two";

    fmt.Printf("Testing F string --- %v --- %v ----", string1, string2);
}

func generateRandomNumberInRange(min int, max int) int8 {
    var tempMax = max - min;
    return int8(rand.Intn(tempMax) + min);
}


