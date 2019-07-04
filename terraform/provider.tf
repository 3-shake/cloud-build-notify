provider "google" {
  version = "2.5"
  project = var.project_id
  region  = var.region
}

provider "archive" {
  version = "~> 1.2"
}
