using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DynamoDB.Table("example", new Aws.DynamoDB.TableArgs
        {
            Attributes = 
            {
                new Aws.DynamoDB.Inputs.TableAttributeArgs
                {
                    Name = "TestTableHashKey",
                    Type = "S",
                },
            },
            BillingMode = "PAY_PER_REQUEST",
            HashKey = "TestTableHashKey",
            Replicas = 
            {
                new Aws.DynamoDB.Inputs.TableReplicaArgs
                {
                    RegionName = "us-east-2",
                },
                new Aws.DynamoDB.Inputs.TableReplicaArgs
                {
                    RegionName = "us-west-2",
                },
            },
            StreamEnabled = true,
            StreamViewType = "NEW_AND_OLD_IMAGES",
        });
    }

}

