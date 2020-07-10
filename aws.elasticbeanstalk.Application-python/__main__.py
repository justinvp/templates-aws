import pulumi
import pulumi_aws as aws

tftest = aws.elasticbeanstalk.Application("tftest",
    appversion_lifecycle={
        "deleteSourceFromS3": True,
        "maxCount": 128,
        "service_role": aws_iam_role["beanstalk_service"]["arn"],
    },
    description="tf-test-desc")

