import pulumi
import pulumi_aws as aws

available = aws.get_availability_zones(state="available")
primary = aws.ec2.Subnet("primary", availability_zone=available.names[0])
# ...
secondary = aws.ec2.Subnet("secondary", availability_zone=available.names[1])
# ...

