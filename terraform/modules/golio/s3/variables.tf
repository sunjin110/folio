variable "name" {
  type = string
}

variable "cors" {
  type = object({
    allowed_origins : list(string)
  })
}
