resource "google_cloud_run_service" "default" {
  name     = "graphql"
  location = var.region_tokyo

  template {
    spec {
      container_concurrency = 80
      timeout_seconds       = 30
      containers {
        image = "gcr.io/golang-portfolio/go_pf/graphql@sha256:292a98ceedafad8beb8fca9f9a1f1c5108a9dc73c258b43eed14bd64a35ed39e"
        env {
          name = "GRAPHQL_SERVICE_PORT"
          value = "80"
        }
        env {
          name = "USER_SERVICE_NAME"
          value = "user"
        }
        env {
          name = "USER_SERVICE_PORT"
          value = "50051"
        }
        resources {
          limits = {
            "cpu":"1000m",
            "memory" : "128Mi"
          }
        }
        ports{
          container_port=80
        }
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location    = google_cloud_run_service.default.location
  project     = google_cloud_run_service.default.project
  service     = google_cloud_run_service.default.name
  policy_data = data.google_iam_policy.noauth.policy_data
}