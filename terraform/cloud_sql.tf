resource "google_sql_database_instance" "instance" {
  name   = "go-pf-instance"
  region = "asia-northeast1"
  settings {
    tier = "db-f1-micro"
  }
}

resource "google_sql_database" "database" {
  name     = "go-pf-db"
  instance = google_sql_database_instance.instance.name
}

resource "google_sql_user" "users" {
  name     = "user"
  instance = google_sql_database_instance.instance.name
  password = "password"
}