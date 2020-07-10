import pulumi
import pulumi_aws as aws

model = aws.sagemaker.Model("model",
    execution_role_arn=aws_iam_role["foo"]["arn"],
    primary_container={
        "image": "174872318107.dkr.ecr.us-west-2.amazonaws.com/kmeans:1",
    })
assume_role = aws.iam.get_policy_document(statements=[{
    "actions": ["sts:AssumeRole"],
    "principals": [{
        "identifiers": ["sagemaker.amazonaws.com"],
        "type": "Service",
    }],
}])
role = aws.iam.Role("role", assume_role_policy=assume_role.json)

