import pulumi
import pulumi_aws as aws

available = aws.get_availability_zones()
current_region = aws.get_region()
current_caller_identity = aws.get_caller_identity()
foo_eip = aws.ec2.Eip("fooEip", vpc=True)
foo_protection = aws.shield.Protection("fooProtection", resource_arn=foo_eip.id.apply(lambda id: f"arn:aws:ec2:{current_region.name}:{current_caller_identity.account_id}:eip-allocation/{id}"))

