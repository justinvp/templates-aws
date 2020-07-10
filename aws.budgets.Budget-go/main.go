package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/budgets"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := budgets.NewBudget(ctx, "ec2", &budgets.BudgetArgs{
			BudgetType: pulumi.String("COST"),
			CostFilters: pulumi.StringMap{
				"Service": pulumi.String("Amazon Elastic Compute Cloud - Compute"),
			},
			LimitAmount: pulumi.String("1200"),
			LimitUnit:   pulumi.String("USD"),
			Notifications: budgets.BudgetNotificationArray{
				&budgets.BudgetNotificationArgs{
					ComparisonOperator: pulumi.String("GREATER_THAN"),
					NotificationType:   pulumi.String("FORECASTED"),
					SubscriberEmailAddresses: pulumi.StringArray{
						pulumi.String("test@example.com"),
					},
					Threshold:     pulumi.Float64(100),
					ThresholdType: pulumi.String("PERCENTAGE"),
				},
			},
			TimePeriodEnd:   pulumi.String("2087-06-15_00:00"),
			TimePeriodStart: pulumi.String("2017-07-01_00:00"),
			TimeUnit:        pulumi.String("MONTHLY"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

