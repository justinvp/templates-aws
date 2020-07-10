import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
current = aws.get_region()
example_product_subscription = aws.securityhub.ProductSubscription("exampleProductSubscription", product_arn=f"arn:aws:securityhub:{current.name}:733251395267:product/alertlogic/althreatmanagement",
opts=ResourceOptions(depends_on=["aws_securityhub_account.example"]))

