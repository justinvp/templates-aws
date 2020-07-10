import pulumi
import pulumi_aws as aws

elasticbeanstalk = aws.iam.ServiceLinkedRole("elasticbeanstalk", aws_service_name="elasticbeanstalk.amazonaws.com")

