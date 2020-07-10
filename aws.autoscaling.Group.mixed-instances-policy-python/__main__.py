import pulumi
import pulumi_aws as aws

example_launch_template = aws.ec2.LaunchTemplate("exampleLaunchTemplate",
    image_id=data["aws_ami"]["example"]["id"],
    instance_type="c5.large",
    name_prefix="example")
example_group = aws.autoscaling.Group("exampleGroup",
    availability_zones=["us-east-1a"],
    desired_capacity=1,
    max_size=1,
    min_size=1,
    mixed_instances_policy={
        "launch_template": {
            "launchTemplateSpecification": {
                "launchTemplateId": example_launch_template.id,
            },
            "override": [
                {
                    "instance_type": "c4.large",
                    "weightedCapacity": "3",
                },
                {
                    "instance_type": "c3.large",
                    "weightedCapacity": "2",
                },
            ],
        },
    })

