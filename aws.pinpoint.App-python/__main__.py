import pulumi
import pulumi_aws as aws

example = aws.pinpoint.App("example",
    limits={
        "maximumDuration": 600,
    },
    quiet_time={
        "end": "06:00",
        "start": "00:00",
    })

