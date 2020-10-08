resource "google_sql_database_instance" "instance" {
  name   = "go-pf"
  region = "asia-northeast1"
  settings {
    tier = "db-f1-micro"
  }
}

resource "google_sql_database" "database" {
  name     = var.db_name
  instance = google_sql_database_instance.instance.name
}

resource "google_sql_user" "users" {
  name     = "user"
  instance = google_sql_database_instance.instance.name
  password = "password"
}