import pulumi
import pulumi_aws as aws

# Create a new load balancer attachment
baz = aws.elb.Attachment("baz",
    elb=aws_elb["bar"]["id"],
    instance=aws_instance["foo"]["id"])

