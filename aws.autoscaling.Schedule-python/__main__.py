import pulumi
import pulumi_aws as aws

foobar_group = aws.autoscaling.Group("foobarGroup",
    availability_zones=["us-west-2a"],
    force_delete=True,
    health_check_grace_period=300,
    health_check_type="ELB",
    max_size=1,
    min_size=1,
    termination_policies=["OldestInstance"])
foobar_schedule = aws.autoscaling.Schedule("foobarSchedule",
    autoscaling_group_name=foobar_group.name,
    desired_capacity=0,
    end_time="2016-12-12T06:00:00Z",
    max_size=1,
    min_size=0,
    scheduled_action_name="foobar",
    start_time="2016-12-11T18:00:00Z")

