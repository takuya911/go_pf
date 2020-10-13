# よく変更する部分
variable "graphql_image" {
  default = "gcr.io/golang-portfolio/go_pf/graphql@sha256:d89b44fe2c602dc98c51e2a0d3ba5395397e2398b5d5cf0dc5ffe1b71193f906"
}

variable "user_image"{
  default="gcr.io/golang-portfolio/go_pf/user@sha256:85edd9b3fed1865cd5a289817e78f30cf0d40ed8b400f08dd7aab23d339abd21"
}

variable "db_instance"{
  default="gopf"
}

variable "db_connect_name"{
  default="golang-portfolio:asia-northeast1:gopf"
}

##################################################
variable "zone_tokyo" {
  default = "asia-northeast1-a"
}

variable "region_tokyo" {
  default = "asia-northeast1"
}

variable "graphql_service_port" {
  default = "80"
}

variable "user_service_name" {
  default = "user-repftyfivq-an.a.run.app"
}

variable "user_service_port" {
  default = "443"
}

variable "db_name"{
  default="go-pf-db"
}

variable "db_user"{
  default="user"
}

variable "db_pass"{
  default="password"
}

