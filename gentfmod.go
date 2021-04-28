package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var modulename string
	rootmodfiles := []string{"main.tf", "variables.tf", "outputs.tf", "provider.tf"}
	childmodfiles := []string{"main.tf", "variables.tf", "outputs.tf", "README.md"}
	providersource := "hashicorp/aws"
	providerversion := ">= 3.37.0"
	tfawsprovider := "terraform{\n\trequired_providers{\n\t\taws = {\n\t\t\tsource = \"" + providersource + "\" \n\t\t\tversion = \"" + providerversion + "\"\n\t\t}\n\t}\n}"
	exampleresource := "resource \"aws_accessanalyzer_analyzer\" \"accessanalyzer\" {\n\tanalyzer_name = var.accessanalyzername\n}"
	childoutput := "output \"accessanalyzerarn\" {\n\tdescription = \"ARN of the IAM Access Analyzer\"\n\tvalue = aws_accessanalyzer_analyzer.accessanalyzer.arn\n}"
	childvar := "variable \"accessanalyzername\" {\n\tdescription = \"Name of the IAM Access Analyzer\"\n\ttype = string\n}"
	rootoutput := "output \"accessanalyzerarnuseast1\" {\n\tdescription = \"ARN of the access analyzer\"\n\tvalue = module.accessanalyzeruseast1.accessanalyzerarn\n}"
	rootproviders := "provider \"aws\" {\n\tregion = \"us-east-1\"\n\talias = \"useast1\"\n\n\tassume_role{\n\t\trole_arn = \"<role arn>\"\n\t\tsession_name = \"<session name>\"\n\t}\n}"

	fmt.Println("Enter terraform module name: ")
	fmt.Scanln(&modulename) //terraform-<PROVIDER>-<NAME>
	fmt.Printf("Terraform module name: %v\n", "terraform-aws-"+modulename)
	rootmain := "module \"accessanalyzeruseast1\" {\n\tsource = \"./modules/" + "terraform-aws-" + modulename + "\"\n\taccessanalyzername = \"accessanalyzeruseast1\"\n\tproviders = {\n\t\taws = aws.useast1\n\t}\n}"

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
	openFileandWrite(d+"/tf"+"/modules"+"/"+"terraform-aws-"+modulename+"/main.tf", tfawsprovider+"\n\n"+exampleresource)
	openFileandWrite(d+"/tf"+"/modules"+"/"+"terraform-aws-"+modulename+"/outputs.tf", childoutput)
	openFileandWrite(d+"/tf"+"/modules"+"/"+"terraform-aws-"+modulename+"/variables.tf", childvar)
	openFileandWrite(d+"/tf"+"/main.tf", rootmain)
	openFileandWrite(d+"/tf"+"/outputs.tf", rootoutput)
	openFileandWrite(d+"/tf"+"/provider.tf", rootproviders)
	fmt.Println("Successfully created folders and generated files to write structured Terraform modules following best practices.")
}

func openFileandWrite(path string, s string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("File could not be opened")
	}

	_, err = f.Write([]byte(s))
	if err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	f.Close()
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
