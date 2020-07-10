import pulumi
import pulumi_aws as aws

example = aws.ec2.Fleet("example",
    launch_template_config={
        "launchTemplateSpecification": {
            "launchTemplateId": aws_launch_template["example"]["id"],
            "version": aws_launch_template["example"]["latest_version"],
        },
    },
    target_capacity_specification={
        "defaultTargetCapacityType": "spot",
        "totalTargetCapacity": 5,
    })

