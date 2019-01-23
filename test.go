package main

func myFunc(){
	println("myFunc Called")
}

func main() {
	var tmpVal = 10
	println("hello!")
	println("hello 2!")
	tmpVal = 11
	myFunc()
	println("hello 3!" , tmpVal)
}
