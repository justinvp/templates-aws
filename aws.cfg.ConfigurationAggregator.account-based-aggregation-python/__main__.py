import pulumi
import pulumi_aws as aws

account = aws.cfg.ConfigurationAggregator("account", account_aggregation_source={
    "accountIds": ["123456789012"],
    "regions": ["us-west-2"],
})

