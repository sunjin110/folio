output "acm" {
  value = {
    certificate_arn = aws_acm_certificate.cert.arn
  }
}