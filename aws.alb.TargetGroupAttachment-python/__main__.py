import pulumi
import pulumi_aws as aws

test_target_group = aws.lb.TargetGroup("testTargetGroup")
test_instance = aws.ec2.Instance("testInstance")
test_target_group_attachment = aws.lb.TargetGroupAttachment("testTargetGroupAttachment",
    port=80,
    target_group_arn=test_target_group.arn,
    target_id=test_instance.id)

