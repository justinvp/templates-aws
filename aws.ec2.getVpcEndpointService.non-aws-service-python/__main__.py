import pulumi
import pulumi_aws as aws

custome = aws.ec2.get_vpc_endpoint_service(service_name="com.amazonaws.vpce.us-west-2.vpce-svc-0e87519c997c63cd8")

