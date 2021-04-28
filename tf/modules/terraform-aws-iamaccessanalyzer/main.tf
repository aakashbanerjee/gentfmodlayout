terraform{
	required_providers{
		aws = {
			source = "hashicorp/aws" 
			version = ">= 3.37.0"
		}
	}
}

resource "aws_accessanalyzer_analyzer" "accessanalyzer" {
	analyzer_name = var.accessanalyzername
}