import pulumi
import pulumi_aws as aws

test = aws.ec2.get_launch_template(filters=[{
    "name": "launch-template-name",
    "values": ["some-template"],
}])

