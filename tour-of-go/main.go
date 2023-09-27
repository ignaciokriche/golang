// Tour of go code and exercises by Ignacio Krichevsky

package main // ran go mod init github.com/ignaciokriche/golang

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("Tour of Go code & exercises by Ignacio Krichevsky")

	// add an extra option with all the options.
	// done here to avoid initialization cycle.
	actions = append(actions, action{"run all", runAll})

	for option, err := readUserOption(); err == nil && option >= 0 && option < len(actions); option, err = readUserOption() {
		fmt.Println("\nyour selection is:", actions[option])
		fmt.Println()
		time.Sleep(1 * time.Second)
		actions[option].demo()
		time.Sleep(2 * time.Second)
	}

}

type action struct {
	text string
	demo func()
}

func (a action) String() string {
	return a.text
}

var actions = []action{
	{"the very basics", ElementalDemo},
	{"iota", IotaDemo},
	{"switch", SwitchDemo},
	{"defer", DeferDemo},
	{"struct", StructDemo},
	{"arrays & slices", ArraySliceDemo},
	{"range", RangeDemo},
	{"picture generator exercise", ShowPicDemo},
	{"maps", MapsDemo},
	{"maps - word count", RunMapsExercice},
	{"functions", FunctionsDemo},
	{"Fibonacci closure exercise", FibonacciDemo},
	{"methods", MethodsDemo},
	{"interface", InterfaceDemo},
	{"type assertion", TypeAssertionsDemo},
	{"stringer exercise", StringerExercise},
	{"error exercise", ErrorsExercise},
	{"reader", ReaderDemo},
	{"rot13 reader exercise", ExerciseRot13Reader},
	{"image exercise", ImageExercise},
	{"type parameters", GenericsDemo},
	{"simple tree exercise", SimpleTreeExercisesDemo},
	{"goroutines", GoroutinesDemo},
	{"Fibonacci channel range close exercise", ChannelRangeCloseDemo},
	{"Fibonacci channel select exercise", SelectDemo},
	{"equivalent einary tree exercise", ExerciseEquivalentBinaryTreeDemo},
	{"web crawler exercise", ExerciseWebCrawlerDemo},
}

func readUserOption() (int, error) {

	var input string
	var err error
	var option int

	fmt.Print("type a number and press enter:\n\n")
	for i, action := range actions {
		fmt.Printf("\t%2d\t%s\n", i, action)
	}
	fmt.Println("------------------------------------------------------")
	fmt.Println()

	if input, err = bufio.NewReader(os.Stdin).ReadString('\n'); err == nil {
		if option, err = strconv.Atoi(strings.TrimSuffix(input, "\n")); err == nil {
			return option, nil
		}
	}
	return 0, err
}

func runAll() {
	for _, a := range actions[:len(actions)-1] {
		a.demo()
	}
}
