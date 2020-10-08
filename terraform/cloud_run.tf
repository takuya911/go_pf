data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

# Graphql Service
resource "google_cloud_run_service" "graphql" {
  name     = "graphql"
  location = var.region_tokyo
  template {
    spec {
      container_concurrency = 80
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
          container_port = 80
        }
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location    = google_cloud_run_service.graphql.location
  project     = google_cloud_run_service.graphql.project
  service     = google_cloud_run_service.graphql.name
  policy_data = data.google_iam_policy.noauth.policy_data
}