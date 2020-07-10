import pulumi
import pulumi_aws as aws

example = aws.macie.S3BucketAssociation("example",
    bucket_name="tf-macie-example",
    classification_type={
        "oneTime": "FULL",
    },
    prefix="data")

