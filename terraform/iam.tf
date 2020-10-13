data "google_iam_policy" "invoker" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "graphql" {
  location    = google_cloud_run_service.graphql.location
  project     = google_cloud_run_service.graphql.project
  service     = google_cloud_run_service.graphql.name
  policy_data = data.google_iam_policy.invoker.policy_data
}

resource "google_cloud_run_service_iam_policy" "user" {
  location    = google_cloud_run_service.user.location
  project     = google_cloud_run_service.user.project
  service     = google_cloud_run_service.user.name
  policy_data = data.google_iam_policy.invoker.policy_data
}