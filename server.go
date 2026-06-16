package main

func main() {
	a := App{}

	a.Initialize("test.db")

	a.Run(":8080")

	print("connection success")
}