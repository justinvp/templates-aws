using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dynamodbTarget = new Aws.AppAutoScaling.Target("dynamodbTarget", new Aws.AppAutoScaling.TargetArgs
        {
            MaxCapacity = 100,
            MinCapacity = 5,
            ResourceId = "table/tableName",
            ScalableDimension = "dynamodb:table:ReadCapacityUnits",
            ServiceNamespace = "dynamodb",
        });
        var dynamodbScheduledAction = new Aws.AppAutoScaling.ScheduledAction("dynamodbScheduledAction", new Aws.AppAutoScaling.ScheduledActionArgs
        {
            ResourceId = dynamodbTarget.ResourceId,
            ScalableDimension = dynamodbTarget.ScalableDimension,
            ScalableTargetAction = new Aws.AppAutoScaling.Inputs.ScheduledActionScalableTargetActionArgs
            {
                MaxCapacity = 200,
                MinCapacity = 1,
            },
            Schedule = "at(2006-01-02T15:04:05)",
            ServiceNamespace = dynamodbTarget.ServiceNamespace,
        });
    }

}

