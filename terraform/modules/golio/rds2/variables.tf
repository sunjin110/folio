variable "prefix" {
  type = string
}

variable "network" {
  type = object({
    vpc_id : string
    private_cidr_blocks : list(string)
    private_subnet_ids : list(string)
  })
}
