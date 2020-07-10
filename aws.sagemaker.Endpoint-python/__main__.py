import pulumi
import pulumi_aws as aws

endpoint = aws.sagemaker.Endpoint("endpoint",
    endpoint_config_name=aws_sagemaker_endpoint_configuration["ec"]["name"],
    tags={
        "Name": "foo",
    })

