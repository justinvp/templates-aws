import pulumi
import pulumi_aws as aws

example = aws.servicequotas.get_service(service_name="Amazon Virtual Private Cloud (Amazon VPC)")

