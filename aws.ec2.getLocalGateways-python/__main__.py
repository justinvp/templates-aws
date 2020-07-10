import pulumi
import pulumi_aws as aws

foo_local_gateways = aws.ec2.get_local_gateways(tags={
    "service": "production",
})
pulumi.export("foo", foo_local_gateways.ids)

