import pulumi
import pulumi_aws as aws

example = aws.gamelift.Alias("example",
    description="Example Description",
    routing_strategy={
        "message": "Example Message",
        "type": "TERMINAL",
    })

