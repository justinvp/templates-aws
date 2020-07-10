import pulumi
import pulumi_aws as aws

my_instance = aws.opsworks.Instance("my-instance",
    instance_type="t2.micro",
    layer_ids=[aws_opsworks_custom_layer["my-layer"]["id"]],
    os="Amazon Linux 2015.09",
    stack_id=aws_opsworks_stack["main"]["id"],
    state="stopped")

