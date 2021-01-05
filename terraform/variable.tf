# よく変更する部分
variable "graphql_image" {
  default = "gcr.io/golang-portfolio/gqlgen-grpc/graphql@sha256:974efa899b8ea00c9891a4309f93f8aa2fcb0021618f2dbcd8ca4ea033fa894a"
}

variable "user_image"{
  default="gcr.io/golang-portfolio/gqlgen-grpc/user@sha256:85edd9b3fed1865cd5a289817e78f30cf0d40ed8b400f08dd7aab23d339abd21"
}

variable "db_instance"{
  default="go-ptfo"
}

variable "db_connect_name"{
  default="golang-portfolio:asia-northeast1:go-ptfo"
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
  default="gqlgen-grpc-db"
}

variable "db_user"{
  default="user"
}

variable "db_pass"{
  default="password"
}