using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ec2 = new Aws.Budgets.Budget("ec2", new Aws.Budgets.BudgetArgs
        {
            BudgetType = "COST",
            CostFilters = 
            {
                { "Service", "Amazon Elastic Compute Cloud - Compute" },
            },
            LimitAmount = "1200",
            LimitUnit = "USD",
            Notifications = 
            {
                new Aws.Budgets.Inputs.BudgetNotificationArgs
                {
                    ComparisonOperator = "GREATER_THAN",
                    NotificationType = "FORECASTED",
                    SubscriberEmailAddresses = 
                    {
                        "test@example.com",
                    },
                    Threshold = 100,
                    ThresholdType = "PERCENTAGE",
                },
            },
            TimePeriodEnd = "2087-06-15_00:00",
            TimePeriodStart = "2017-07-01_00:00",
            TimeUnit = "MONTHLY",
        });
    }

}

