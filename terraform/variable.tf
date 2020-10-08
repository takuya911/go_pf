variable "zone_tokyo" {
  default = "asia-northeast1-a"
}

variable "region_tokyo"{
  default="asia-northeast1"
}

variable "graphql_image"{
  default="gcr.io/golang-portfolio/go_pf/graphql@sha256:292a98ceedafad8beb8fca9f9a1f1c5108a9dc73c258b43eed14bd64a35ed39e"
}

variable "graphql_service_port"{
  default="80"
}

variable "user_service_name"{
  default="user"
}

variable "user_service_port"{
  default="80"
}

