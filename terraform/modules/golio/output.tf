
output "golio" {
  value = {
    base_url = module.api_gateway.api_gateway.invoke_url
  }
}
