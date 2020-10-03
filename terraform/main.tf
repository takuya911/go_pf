terraform {
  required_version = "~> 0.12.29"
}

provider "aws" {
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
  region     = "ap-northeast-1"
  version    = "~> 3.0"
}