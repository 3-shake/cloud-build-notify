resource "google_cloudfunctions_function" "cloud_build_notify" {
  name        = var.cloudfunction
  description = ""
  runtime     = "go111"
  region      = "asia-northeast1"

  count                 = 1
  available_memory_mb   = 128
  timeout               = 60
  entry_point           = "NotifyCloudBuild"
  source_archive_bucket = google_storage_bucket.cloud_build_notify.name
  source_archive_object = google_storage_bucket_object.cloud_build_notify.name

  event_trigger {
    event_type = "providers/cloud.pubsub/eventTypes/topic.publish"
    resource   = "cloud-builds"
  }

  environment_variables = {
    SLACK_URL    = var.slack_url
    CHANNEL      = var.channel 
    REPO_NAME    = var.repo_name 
    BRANCH_NAME  = var.branch_name
  }
}
