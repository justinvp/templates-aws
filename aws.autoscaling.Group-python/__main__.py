import pulumi
import pulumi_aws as aws

test = aws.ec2.PlacementGroup("test", strategy="cluster")
bar = aws.autoscaling.Group("bar",
    desired_capacity=4,
    force_delete=True,
    health_check_grace_period=300,
    health_check_type="ELB",
    initial_lifecycle_hooks=[{
        "default_result": "CONTINUE",
        "heartbeat_timeout": 2000,
        "lifecycle_transition": "autoscaling:EC2_INSTANCE_LAUNCHING",
        "name": "foobar",
        "notification_metadata": """{
  "foo": "bar"
}

""",
        "notification_target_arn": "arn:aws:sqs:us-east-1:444455556666:queue1*",
        "role_arn": "arn:aws:iam::123456789012:role/S3Access",
    }],
    launch_configuration=aws_launch_configuration["foobar"]["name"],
    max_size=5,
    min_size=2,
    placement_group=test.id,
    tags=[
        {
            "key": "foo",
            "propagateAtLaunch": True,
            "value": "bar",
        },
        {
            "key": "lorem",
            "propagateAtLaunch": False,
            "value": "ipsum",
        },
    ],
    vpc_zone_identifiers=[
        aws_subnet["example1"]["id"],
        aws_subnet["example2"]["id"],
    ])

