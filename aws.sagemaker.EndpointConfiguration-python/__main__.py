import pulumi
import pulumi_aws as aws

ec = aws.sagemaker.EndpointConfiguration("ec",
    production_variants=[{
        "initialInstanceCount": 1,
        "instance_type": "ml.t2.medium",
        "modelName": aws_sagemaker_model["m"]["name"],
        "variantName": "variant-1",
    }],
    tags={
        "Name": "foo",
    })

