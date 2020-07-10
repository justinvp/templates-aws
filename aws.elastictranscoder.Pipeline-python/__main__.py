import pulumi
import pulumi_aws as aws

bar = aws.elastictranscoder.Pipeline("bar",
    content_config={
        "bucket": aws_s3_bucket["content_bucket"]["bucket"],
        "storage_class": "Standard",
    },
    input_bucket=aws_s3_bucket["input_bucket"]["bucket"],
    role=aws_iam_role["test_role"]["arn"],
    thumbnail_config={
        "bucket": aws_s3_bucket["thumb_bucket"]["bucket"],
        "storage_class": "Standard",
    })

