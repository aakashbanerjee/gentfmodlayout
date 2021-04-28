package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var modulename string
	rootmodfiles := []string{"main.tf", "variables.tf", "output.tf", "provider.tf"}
	childmodfiles := []string{"main.tf", "variables.tf", "output.tf"}
	fmt.Println("Enter terraform module name: ")
	fmt.Scanln(&modulename) //terraform-<PROVIDER>-<NAME>
	fmt.Printf("Terraform module name: %v\n", "terraform-aws-"+modulename)
	d, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Creating terraform module %v folders and files at %v\n", modulename, d)
	makeDir("tf")
	changeDir(d + "/tf")
	createFile(rootmodfiles)
	makeDir("modules")
	changeDir(d + "/tf" + "/modules")
	makeDir("terraform-aws-" + modulename)
	changeDir(d + "/tf" + "/modules" + "/" + "terraform-aws-" + modulename)
	createFile(childmodfiles)
}

func createFile(fn []string) {
	for _, f := range fn {
		fi, err := os.Create(f)
		if err != nil {
			log.Fatalln(err)
		}
		defer fi.Close()
	}

}

func changeDir(dirname string) {
	err := os.Chdir(dirname)
	if err != nil {
		log.Fatalln(err)
	}
}

func makeDir(dirname string) {
	err := os.Mkdir(dirname, 0777)
	if err != nil {
		log.Fatalln(err)
	}
}
