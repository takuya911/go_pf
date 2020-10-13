# よく変更する部分
variable "graphql_image" {
  default = "gcr.io/golang-portfolio/go_pf/graphql@sha256:f6019dba79d27677ed8d286fd2a1a760850e6ad34b7837e41e4a657747825914"
}

variable "user_image"{
  default="gcr.io/golang-portfolio/go_pf/user@sha256:85edd9b3fed1865cd5a289817e78f30cf0d40ed8b400f08dd7aab23d339abd21"
}

variable "db_instance"{
  default="gopfi"
}

variable "db_connect_name"{
  default="golang-portfolio:asia-northeast1:gopfi"
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

