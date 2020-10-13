variable "zone_tokyo" {
  default = "asia-northeast1-a"
}

variable "region_tokyo" {
  default = "asia-northeast1"
}

variable "graphql_image" {
  default = "gcr.io/golang-portfolio/go_pf/graphql@sha256:b8ff013a3d111bc374c9597f8d234301ce2af35b3b829b3d5352cf18dade8ef2"
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
  default="gcr.io/golang-portfolio/go_pf/user@sha256:bda031d12e6879a8574236b1313e480c88fda094f2f3d253af8f2254183b66bb"
}

variable "db_instance"{
  default="gopf"
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
  default="golang-portfolio:asia-northeast1:gopf"
}

