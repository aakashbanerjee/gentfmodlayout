# gentfmodlayout
Creates folder structure and files for writing terraform modules
#Folder structure
tf
  modules
    modulename
      main.tf
      variables.tf
      output.tf
  main.tf
  variables.tf
  output.tf
  provider.tf
#generate the executable with go build gentfmod.go
#add the symlink to the executable
#sudo ln -s /Users/username/go/src/github.com/gentfmodlayout/gentfmod /usr/local/bin
#cd into any folder and run the executable gentfmod to create the terraform module folder and files in a moment
