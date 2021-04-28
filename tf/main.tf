module "accessanalyzeruseast1" {
	source = "./modules/terraform-aws-iamaccessanalyzer"
	accessanalyzername = "accessanalyzeruseast1"
	providers = {
		aws = aws.useast1
	}
}