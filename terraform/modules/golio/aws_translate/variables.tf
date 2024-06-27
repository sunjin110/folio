variable "aws" {
  type = object({
    account_id = string,
    region     = string,
    profile    = string,
  })
}

variable "prefix" {
  type = string
}

# variable "network" {
#   type = object({
#     vpc_id : string,
#   })
# }

variable "security_group_ids" {
  type = list(string)
}
