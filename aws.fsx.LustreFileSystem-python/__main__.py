import pulumi
import pulumi_aws as aws

example = aws.fsx.LustreFileSystem("example",
    import_path=f"s3://{aws_s3_bucket['example']['bucket']}",
    storage_capacity=1200,
    subnet_ids=aws_subnet["example"]["id"])

