using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dynamodbTableReadTarget = new Aws.AppAutoScaling.Target("dynamodbTableReadTarget", new Aws.AppAutoScaling.TargetArgs
        {
            MaxCapacity = 100,
            MinCapacity = 5,
            ResourceId = "table/tableName",
            ScalableDimension = "dynamodb:table:ReadCapacityUnits",
            ServiceNamespace = "dynamodb",
        });
        var dynamodbTableReadPolicy = new Aws.AppAutoScaling.Policy("dynamodbTableReadPolicy", new Aws.AppAutoScaling.PolicyArgs
        {
            PolicyType = "TargetTrackingScaling",
            ResourceId = dynamodbTableReadTarget.ResourceId,
            ScalableDimension = dynamodbTableReadTarget.ScalableDimension,
            ServiceNamespace = dynamodbTableReadTarget.ServiceNamespace,
            TargetTrackingScalingPolicyConfiguration = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationArgs
            {
                PredefinedMetricSpecification = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs
                {
                    PredefinedMetricType = "DynamoDBReadCapacityUtilization",
                },
                TargetValue = 70,
            },
        });
    }

}

