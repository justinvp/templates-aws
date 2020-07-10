import pulumi
import pulumi_aws as aws

foobar_group = aws.autoscaling.Group("foobarGroup",
    availability_zones=["us-west-2a"],
    health_check_type="EC2",
    tags=[{
        "key": "Foo",
        "propagateAtLaunch": True,
        "value": "foo-bar",
    }],
    termination_policies=["OldestInstance"])
foobar_lifecycle_hook = aws.autoscaling.LifecycleHook("foobarLifecycleHook",
    autoscaling_group_name=foobar_group.name,
    default_result="CONTINUE",
    heartbeat_timeout=2000,
    lifecycle_transition="autoscaling:EC2_INSTANCE_LAUNCHING",
    notification_metadata="""{
  "foo": "bar"
}

""",
    notification_target_arn="arn:aws:sqs:us-east-1:444455556666:queue1*",
    role_arn="arn:aws:iam::123456789012:role/S3Access")

