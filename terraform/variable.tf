variable "zone_tokyo" {
  default = "asia-northeast1-a"
}

variable "region_tokyo" {
  default = "asia-northeast1"
}

variable "graphql_image" {
  default = "gcr.io/golang-portfolio/go_pf/graphql@sha256:292a98ceedafad8beb8fca9f9a1f1c5108a9dc73c258b43eed14bd64a35ed39e"
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

variable "user_image"{
  default="gcr.io/golang-portfolio/go_pf/user@sha256:5bb6ce1e00cdf4711d05657e05d915f5a43e01ec7b9c4fb66da3b317db4027e9"
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

variable "db_connect_name"{
  default="golang-portfolio:asia-northeast1:go-pf-instance"
}