import pulumi
import pulumi_aws as aws

example = aws.route53.get_resolver_rules(tags=[{
    "Environment": "dev",
}])

