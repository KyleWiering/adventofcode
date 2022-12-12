package main

import (
	"fmt"
	"math"
	//"math/big"
)

var lcm = 0
var isFirstCase = false
var isDebug = false
var isDummyPayload = false

type item struct {
	Target int
	Value  float64
}

type monkey struct {
	Inpections float64
	Items      []float64
	Operation  func(float64) float64
	Test       func(float64) float64
}

func dummyInput() []monkey {
	monkeys := []monkey{monkey{ //0
		Inpections: 0,
		Items:      []float64{79, 98},
		Operation:  func(old float64) float64 { return old * float64(19) },
		Test: func(val float64) float64 {
			if math.Remainder(val, 23) == 0 {
				return 2
			} else {
				return 3
			}
		},
	}, monkey{ //1
		Inpections: 0,
		Items:      []float64{54, 65, 75, 74},
		Operation:  func(old float64) float64 { return old + float64(6) },
		Test: func(val float64) float64 {
			if math.Remainder(val, 19) == 0 {
				return 2
			} else {
				return 0
			}
		},
	}, monkey{ //2
		Inpections: 0,
		Items:      []float64{79, 60, 97},
		Operation:  func(old float64) float64 { return old * old },
		Test: func(val float64) float64 {
			if math.Remainder(val, 13) == 0 {
				return 1
			} else {
				return 3
			}
		},
	}, monkey{ //3
		Inpections: 0,
		Items:      []float64{74},
		Operation:  func(old float64) float64 { return old + float64(3) },
		Test: func(val float64) float64 {
			if math.Remainder(val, 17) == 0 {
				return 0
			} else {
				return 1
			}
		},
	},
	}

	return monkeys
}

func realinput() []monkey {

	monkeys := []monkey{monkey{ // 0
		Inpections: 0,
		Items:      []float64{52, 60, 85, 69, 75, 75},
		Operation:  func(old float64) float64 { return old * float64(17) },
		Test: func(val float64) float64 {
			if math.Remainder(val, 13) == 0 {
				return 6
			} else {
				return 7
			}
		},
	}, monkey{ //1
		Inpections: 0,
		Items:      []float64{96, 82, 61, 99, 82, 84, 85},
		Operation:  func(old float64) float64 { return old + float64(8) },
		Test: func(val float64) float64 {
			if math.Remainder(val, 7) == 0 {
				return 0
			} else {
				return 7
			}
		},
	}, monkey{ //2
		Inpections: 0,
		Items:      []float64{95, 79},
		Operation:  func(old float64) float64 { return old + float64(6) },
		Test: func(val float64) float64 {
			if math.Remainder(val, 19) == 0 {
				return 5
			} else {
				return 3
			}
		},
	}, monkey{ // 3
		Inpections: 0,
		Items:      []float64{88, 50, 82, 65, 77},
		Operation:  func(old float64) float64 { return old * float64(19) },
		Test: func(val float64) float64 {
			if math.Remainder(val, 2) == 0 {
				return 4
			} else {
				return 1
			}
		},
	}, monkey{ // 4
		Inpections: 0,
		Items:      []float64{66, 90, 59, 90, 87, 63, 53, 88},
		Operation:  func(old float64) float64 { return old + float64(7) },
		Test: func(val float64) float64 {
			if math.Remainder(val, 5) == 0 {
				return 1
			} else {
				return 0
			}
		},
	}, monkey{ //5
		Inpections: 0,
		Items:      []float64{92, 75, 62},
		Operation:  func(old float64) float64 { return old * old },
		Test: func(val float64) float64 {
			if math.Remainder(val, 3) == 0 {
				return 3
			} else {
				return 4
			}
		},
	}, monkey{ //6
		Inpections: 0,
		Items:      []float64{94, 86, 76, 67},
		Operation:  func(old float64) float64 { return old + float64(1) },
		Test: func(val float64) float64 {
			if math.Remainder(val, 11) == 0 {
				return 5
			} else {
				return 2
			}
		},
	}, monkey{ //7
		Inpections: 0,
		Items:      []float64{57},
		Operation:  func(old float64) float64 { return old + float64(2) },
		Test: func(val float64) float64 {
			if math.Remainder(val, 17) == 0 {
				return 6
			} else {
				return 2
			}
		},
	},
	}

	return monkeys
}

func main() {

	run()
}

func run() {
	var input []monkey

	if isDummyPayload {
		input = dummyInput()
		lcm = LCM(23, 19, 13, 17)
	} else {
		input = realinput()
		lcm = LCM(13, 7, 19, 2, 5, 3, 11, 17)
	}

	fmt.Println("lcm", lcm)
	FirstCase(input)

}

func FirstCase(monkeys []monkey) {
	for i := 0; i < 10000; i++ {
		if i%1000 == 0 {
			fmt.Println("== After Round", i, "==")
			for k, printer := range monkeys {
				fmt.Println("monkey", k, "inspected times", printer.Inpections, "times")
			}
		}

		for j, Monkey := range monkeys {
			if isDebug {
				fmt.Println("monkey", j)
			}

			flyer := ProcessMonkeyItems(&Monkey)
			monkeys[j].Inpections = monkeys[j].Inpections + float64(len(monkeys[j].Items))
			monkeys[j].Items = []float64{}

			for _, fly := range flyer {
				monkeys[fly.Target].Items = append(monkeys[fly.Target].Items, fly.Value)
			}
		}
	}

	first := float64(0)
	second := float64(0)

	for _, Monkey := range monkeys {
		if Monkey.Inpections > first {
			second = first
			first = Monkey.Inpections
		}
	}
	fmt.Println(first, second)
	fmt.Printf("%f", first*second)
	fmt.Println()
}

func ProcessMonkeyItems(Monkey *monkey) []item {
	var tossable []item
	for _, thing := range Monkey.Items {
		if isDebug {
			fmt.Println("\tMonkey inspects an item with a worry level of", thing)
		}
		newVal := Monkey.Operation(thing)
		if isDebug {
			fmt.Println("\t\tWorry level is multiplied by amount to ", newVal)
		}

		if isFirstCase {
			newVal = math.Floor(newVal / float64(3))
		} else {
			newVal = math.Remainder(newVal, float64(lcm))
		}
		if isDebug {
			fmt.Println("\t\tMonkey gets bored with item. Worry level is divided by 3 to ", newVal)
		}
		newTarget := Monkey.Test(newVal)
		if isDebug {
			fmt.Println("\t\tItem with worry level", newVal, "is thrown to", newTarget)
		}
		tossable = append(tossable, item{Target: int(newTarget), Value: newVal})
	}

	return tossable //, Monkey.Inpections
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
