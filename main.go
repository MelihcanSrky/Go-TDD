package main

func main() {
	app := App{}
	app.Initialize("postgres", "admin", "postgres")

	app.Run(":8080")
}
