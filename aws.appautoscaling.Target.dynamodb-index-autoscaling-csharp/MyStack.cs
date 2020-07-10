using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dynamodbIndexReadTarget = new Aws.AppAutoScaling.Target("dynamodbIndexReadTarget", new Aws.AppAutoScaling.TargetArgs
        {
            MaxCapacity = 100,
            MinCapacity = 5,
            ResourceId = $"table/{aws_dynamodb_table.Example.Name}/index/{@var.Index_name}",
            ScalableDimension = "dynamodb:index:ReadCapacityUnits",
            ServiceNamespace = "dynamodb",
        });
    }

}

