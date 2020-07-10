import pulumi
import pulumi_aws as aws

example = aws.datasync.Task("example",
    destination_location_arn=aws_datasync_location_s3["destination"]["arn"],
    options={
        "bytesPerSecond": -1,
    },
    source_location_arn=aws_datasync_location_nfs["source"]["arn"])

