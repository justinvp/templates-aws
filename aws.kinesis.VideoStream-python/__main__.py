import pulumi
import pulumi_aws as aws

default = aws.kinesis.VideoStream("default",
    data_retention_in_hours=1,
    device_name="kinesis-video-device-name",
    media_type="video/h264",
    tags={
        "Name": "kinesis-video-stream",
    })

