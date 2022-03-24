package main

import "go-database/internal"

func main() {
	conf := internal.NewDefaultConfig()
	internal.NewApp(conf).Run()
}
