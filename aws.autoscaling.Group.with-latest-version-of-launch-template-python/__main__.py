import pulumi
import pulumi_aws as aws

foobar = aws.ec2.LaunchTemplate("foobar",
    image_id="ami-1a2b3c",
    instance_type="t2.micro",
    name_prefix="foobar")
bar = aws.autoscaling.Group("bar",
    availability_zones=["us-east-1a"],
    desired_capacity=1,
    launch_template={
        "id": foobar.id,
        "version": "$Latest",
    },
    max_size=1,
    min_size=1)

