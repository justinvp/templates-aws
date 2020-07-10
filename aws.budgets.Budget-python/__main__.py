import pulumi
import pulumi_aws as aws

ec2 = aws.budgets.Budget("ec2",
    budget_type="COST",
    cost_filters={
        "Service": "Amazon Elastic Compute Cloud - Compute",
    },
    limit_amount="1200",
    limit_unit="USD",
    notifications=[{
        "comparison_operator": "GREATER_THAN",
        "notification_type": "FORECASTED",
        "subscriberEmailAddresses": ["test@example.com"],
        "threshold": 100,
        "thresholdType": "PERCENTAGE",
    }],
    time_period_end="2087-06-15_00:00",
    time_period_start="2017-07-01_00:00",
    time_unit="MONTHLY")

