package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Range [2]int

type TypeToAdd struct {
    Letters Range
    Numbers Range
    Symbols Range
}

var typeToAdd = TypeToAdd {
    Letters: Range{97, 122},
    Numbers: Range{48, 57},
    Symbols: Range{33, 47},
}

func main() {
    var numberOfLetters, lettersErr = takeUserInput("Enter number of letters");
    var password string = "";

    if handleError(lettersErr) {
        return;
    }

    password += addToPassword(numberOfLetters, typeToAdd.Letters);

    var numberOfNumbers, numbersErr = takeUserInput("Enter number of numbers");

    if handleError(numbersErr) {
        return;
    }

    password += addToPassword(numberOfNumbers, typeToAdd.Numbers);

    var numberOfSymbols, symbolsErr = takeUserInput("Enter number of symbols");

    if handleError(symbolsErr) {
        return;
    }

    password += addToPassword(numberOfSymbols, typeToAdd.Symbols);

    println(shufflePassword(password));
}

func shufflePassword(password string) string {
    var chars = []rune(password);
    var charsLen = len(chars) - 1;

    for i := 0; i <= charsLen; i++ {
        var randomI = generateRandomNumberInRange(0, charsLen)
        var temp = chars[i];
        chars[i] = chars[randomI];
        chars[randomI] = temp;
    }

    return string(chars);
}

func addToPassword(numberOfChars int, typeToAdd Range) string {
    var result = "";
    var min, max = typeToAdd[0], typeToAdd[1];

    for range numberOfChars {
        result += string(generateRandomNumberInRange(min, max));
    }

    return result;
}

func takeUserInput(msg string) (int, error) {
    var reader = bufio.NewReader(os.Stdin);

    println(msg);
    var input, err = reader.ReadString('\n');

    input = strings.TrimSpace(input);

    if err != nil {
        return 0, err;
    }

    var toNumber, convertError = strconv.Atoi(input);

    if handleError(convertError) {
        return 0, convertError;
    }

    return toNumber, err;
}

func handleError(err error) bool {
    if err != nil {
        println(err.Error());

        return true;
    }

    return false;
}

func generateRandomNumberInRange(min int, max int) int8 {
    var tempMax = max - min;

    return int8(rand.Intn(tempMax) + min);
}
