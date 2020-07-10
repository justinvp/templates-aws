import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ec2 = new aws.budgets.Budget("ec2", {
    budgetType: "COST",
    costFilters: {
        Service: "Amazon Elastic Compute Cloud - Compute",
    },
    limitAmount: "1200",
    limitUnit: "USD",
    notifications: [{
        comparisonOperator: "GREATER_THAN",
        notificationType: "FORECASTED",
        subscriberEmailAddresses: ["test@example.com"],
        threshold: 100,
        thresholdType: "PERCENTAGE",
    }],
    timePeriodEnd: "2087-06-15_00:00",
    timePeriodStart: "2017-07-01_00:00",
    timeUnit: "MONTHLY",
});

