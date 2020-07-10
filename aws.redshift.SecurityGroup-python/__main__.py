import pulumi
import pulumi_aws as aws

default = aws.redshift.SecurityGroup("default", ingress=[{
    "cidr": "10.0.0.0/24",
}])

