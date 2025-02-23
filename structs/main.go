package main

type gasEgnine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
    mpkwh uint8;
    kwh uint8;
}

type engine interface {
    milesLeft() uint8;
}

func main() {
	var engine gasEgnine = gasEgnine{25, 60};
	var elEngine electricEngine = electricEngine{25, 60};

    println(canMakeIt(engine, 100));
    println(canMakeIt(elEngine, 100));
}

func (engine gasEgnine) milesLeft() uint8 {
    return engine.gallons * engine.mpg;
}

func (engine electricEngine) milesLeft() uint8 {
    return engine.kwh * engine.mpkwh;
}

func canMakeIt(engine engine, miles uint8) bool {
	return engine.milesLeft() > miles;
}
