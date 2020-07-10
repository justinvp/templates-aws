import pulumi
import pulumi_aws as aws

console = aws.cloudwatch.EventRule("console",
    description="Capture all EC2 scaling events",
    event_pattern="""{
  "source": [
    "aws.autoscaling"
  ],
  "detail-type": [
    "EC2 Instance Launch Successful",
    "EC2 Instance Terminate Successful",
    "EC2 Instance Launch Unsuccessful",
    "EC2 Instance Terminate Unsuccessful"
  ]
}

""")
test_stream = aws.kinesis.Stream("testStream", shard_count=1)
yada = aws.cloudwatch.EventTarget("yada",
    arn=test_stream.arn,
    rule=console.name,
    run_command_targets=[
        {
            "key": "tag:Name",
            "values": ["FooBar"],
        },
        {
            "key": "InstanceIds",
            "values": ["i-162058cd308bffec2"],
        },
    ])

