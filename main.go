package main

import (
	"log"

	functionn "main.go/functions"
	controller "main.go/newcontroller"
)

func main() {
	uuid := "8433f2c4-6446-4994-b66d-86e67c81e056"
	commonlogpath := "C://Users//nisha//OneDrive//Desktop//demolog//log (1)//log"
	dir := "info"
	data, err := functionn.Readappdata(uuid, dir, commonlogpath)
	if err != nil {
		log.Fatal("Error in Reading app data")
	}
	controller.Insert_Mongodb(data)
}
