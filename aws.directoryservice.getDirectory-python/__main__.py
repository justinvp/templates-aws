import pulumi
import pulumi_aws as aws

example = aws.directoryservice.get_directory(directory_id=aws_directory_service_directory["main"]["id"])

