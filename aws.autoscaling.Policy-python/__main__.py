import pulumi
import pulumi_aws as aws

bar = aws.autoscaling.Group("bar",
    availability_zones=["us-east-1a"],
    force_delete=True,
    health_check_grace_period=300,
    health_check_type="ELB",
    launch_configuration=aws_launch_configuration["foo"]["name"],
    max_size=5,
    min_size=2)
bat = aws.autoscaling.Policy("bat",
    adjustment_type="ChangeInCapacity",
    autoscaling_group_name=bar.name,
    cooldown=300,
    scaling_adjustment=4)

