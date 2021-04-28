**gentfmodlayout (Tested on MacOS only)**

*Creates folder structure and files for writing terraform modules*

**Folder structure**

```
tf
  modules
    terraform-aws-modulename
      main.tf
      variables.tf
      output.tf
  main.tf
  variables.tf
  output.tf
  provider.tf

```

generate the executable with go build gentfmod.go

add the symlink to the executable

`sudo ln -s /Users/username/go/src/github.com/gentfmodlayout/gentfmod /usr/local/bin`

cd into any folder and run the executable gentfmod to create the terraform module folder and files in a moment

setup terraform workspace to track changes to /tf and you arre all set for a smooth IAC project setup.

added example terraform code to create a free resource IAM Access Analyzer in AWS account us-east-1.

**Please note you will need to add the role arn and session name in the root module provider.tf**
