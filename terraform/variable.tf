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

variable "slack_token" {
  default = ""
}

variable "slack_channel_id" {
  default = ""
}
