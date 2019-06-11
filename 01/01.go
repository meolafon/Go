package main

import "fmt"

func main() {
	//hello("Billy")
	a, b := swap("b", "a")
	fmt.Println(a, b)
	//fmt.Println("This was compiled with Go on Windows")
	//a := "asdf"
}

// func hello(name string) {
// 	fmt.Println("hello, " + name + " !")
// }

// x & y are of the same type - declare once
// returns tuple
func swap(x, y string) (string, string) {
	return y, x
}

// func readFile() {
// 	f, err = os.Open("file.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	r := bufio.NewReader(f)
// 	for {
// 		s, err := r.ReadString('\n')
// 	}

//}
