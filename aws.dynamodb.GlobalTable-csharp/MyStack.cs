using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var us_east_1 = new Aws.Provider("us-east-1", new Aws.ProviderArgs
        {
            Region = "us-east-1",
        });
        var us_west_2 = new Aws.Provider("us-west-2", new Aws.ProviderArgs
        {
            Region = "us-west-2",
        });
        var us_east_1Table = new Aws.DynamoDB.Table("us-east-1Table", new Aws.DynamoDB.TableArgs
        {
            Attributes = 
            {
                new Aws.DynamoDB.Inputs.TableAttributeArgs
                {
                    Name = "myAttribute",
                    Type = "S",
                },
            },
            HashKey = "myAttribute",
            ReadCapacity = 1,
            StreamEnabled = true,
            StreamViewType = "NEW_AND_OLD_IMAGES",
            WriteCapacity = 1,
        }, new CustomResourceOptions
        {
            Provider = "aws.us-east-1",
        });
        var us_west_2Table = new Aws.DynamoDB.Table("us-west-2Table", new Aws.DynamoDB.TableArgs
        {
            Attributes = 
            {
                new Aws.DynamoDB.Inputs.TableAttributeArgs
                {
                    Name = "myAttribute",
                    Type = "S",
                },
            },
            HashKey = "myAttribute",
            ReadCapacity = 1,
            StreamEnabled = true,
            StreamViewType = "NEW_AND_OLD_IMAGES",
            WriteCapacity = 1,
        }, new CustomResourceOptions
        {
            Provider = "aws.us-west-2",
        });
        var myTable = new Aws.DynamoDB.GlobalTable("myTable", new Aws.DynamoDB.GlobalTableArgs
        {
            Replicas = 
            {
                new Aws.DynamoDB.Inputs.GlobalTableReplicaArgs
                {
                    RegionName = "us-east-1",
                },
                new Aws.DynamoDB.Inputs.GlobalTableReplicaArgs
                {
                    RegionName = "us-west-2",
                },
            },
        }, new CustomResourceOptions
        {
            Provider = "aws.us-east-1",
            DependsOn = 
            {
                "aws_dynamodb_table.us-east-1",
                "aws_dynamodb_table.us-west-2",
            },
        });
    }

}

