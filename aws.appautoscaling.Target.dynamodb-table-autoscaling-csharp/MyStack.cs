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
            ResourceId = $"table/{aws_dynamodb_table.Example.Name}",
            ScalableDimension = "dynamodb:table:ReadCapacityUnits",
            ServiceNamespace = "dynamodb",
        });
    }

}

