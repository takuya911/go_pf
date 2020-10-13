# Graphql Service
resource "google_cloud_run_service" "graphql" {
  name     = "graphql"
  location = var.region_tokyo
  template {
    spec {
      container_concurrency = 20
      timeout_seconds       = 30
      containers {
        image = var.graphql_image
        env {
          name  = "GRAPHQL_SERVICE_PORT"
          value = var.graphql_service_port
        }
        env {
          name  = "USER_SERVICE_NAME"
          value = var.user_service_name
        }
        env {
          name  = "USER_SERVICE_PORT"
          value = var.user_service_port
        }
        resources {
          limits = {
            "cpu" : "1000m",
            "memory" : "128Mi"
          }
        }
        ports {
          container_port = var.graphql_service_port
        }
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}

# User Service
resource "google_cloud_run_service" "user" {
  name     = "user"
  location = var.region_tokyo
  template {
    metadata {
      annotations = {
        "run.googleapis.com/cloudsql-instances" = var.db_connect_name
      }
    }
    spec {
      container_concurrency = 20
      timeout_seconds       = 30
      containers {
        image = var.user_image
        env {
          name  = "USER_SERVICE_PORT"
          value = var.user_service_port
        }
        env {
          name  = "DB_NAME"
          value = var.db_name
        }
        env {
          name  = "DB_USER"
          value = var.db_user
        }
        env {
          name  = "DB_PASS"
          value = var.db_pass
        }
        env {
          name  = "DB_CONNECT_NAME"
          value = var.db_connect_name
        }
        resources {
          limits = {
            "cpu" : "1000m",
            "memory" : "128Mi"
          }
        }
        ports {
          container_port = var.user_service_port
        }
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}
