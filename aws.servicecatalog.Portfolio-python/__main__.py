import pulumi
import pulumi_aws as aws

portfolio = aws.servicecatalog.Portfolio("portfolio",
    description="List of my organizations apps",
    provider_name="Brett")

