resource "google_storage_bucket" "cloud_build_notify" {
  name          = var.bucket
  location      = var.region
  storage_class = "REGIONAL"

  labels = {
    project = var.project
  }

  versioning {
    enabled = false
  }

  force_destroy = true
}

data "archive_file" "cloud_build_notify" {
  type        = "zip"
  source_dir  = "./cloudfunction/cloud-build-notify"
  output_path = "./cloudfunction/cloud-build-notify.zip"
}

resource "google_storage_bucket_object" "cloud_build_notify" {
  name   = "cloud_build_notify/cloud-build-notify-${timestamp()}.zip"
  source = "${data.archive_file.cloud_build_notify.output_path}"
  bucket = google_storage_bucket.cloud_build_notify.name
}
