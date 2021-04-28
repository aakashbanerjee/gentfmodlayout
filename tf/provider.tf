provider "aws" {
	region = "us-east-1"
	alias = "useast1"

	assume_role{
		role_arn = "<role arn>"
		session_name = "<session name>"
	}
}