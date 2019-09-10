variable "project_id" {
  description = "Project ID / gcloud projects list"
  default     = ""
}

variable "project" {
  description = "Project Name / gcloud projects list"
  default     = ""
}

variable "bucket" {
  description = "Bucket Name / gsutil ls"
  default     = ""
}

variable "cloudfunction" {
  description = "Cloud Function Name / gcloud functions"
  default     = ""
}

variable "location" {
  description = ""
  default     = "taiwan"
}

variable "region" {
  default = "asia-east1"
}

variable "slack_url" {
  default = ""
}

variable "channel" {
  default = ""
}

variable "repo_name" {
  default = "asia-east1"
}

variable "branch_name" {
  default = "asia-east1"
}
