import pulumi
import pulumi_aws as aws

example = aws.gamelift.Fleet("example",
    build_id=aws_gamelift_build["example"]["id"],
    ec2_instance_type="t2.micro",
    fleet_type="ON_DEMAND",
    runtime_configuration={
        "serverProcess": [{
            "concurrentExecutions": 1,
            "launchPath": "C:\\game\\GomokuServer.exe",
        }],
    })

