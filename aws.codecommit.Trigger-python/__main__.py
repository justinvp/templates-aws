import pulumi
import pulumi_aws as aws

test_repository = aws.codecommit.Repository("testRepository", repository_name="test")
test_trigger = aws.codecommit.Trigger("testTrigger",
    repository_name=test_repository.repository_name,
    triggers=[{
        "destination_arn": aws_sns_topic["test"]["arn"],
        "events": ["all"],
        "name": "all",
    }])

