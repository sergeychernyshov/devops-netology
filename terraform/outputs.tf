output "account_id" {
  value = data.aws_caller_identity.current.account_id
}

output "caller_arn" {
  value = data.aws_caller_identity.current.arn
}

output "caller_user" {
  value = data.aws_caller_identity.current.user_id
}

output "instance_ip_addr" {
  value = aws_instance.server.private_ip
}

output "subnet_public_id" {
  value = aws_subnet.aws_subnet_public.id
}

