using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var basic_dynamodb_table = new Aws.DynamoDB.Table("basic-dynamodb-table", new Aws.DynamoDB.TableArgs
        {
            Attributes = 
            {
                new Aws.DynamoDB.Inputs.TableAttributeArgs
                {
                    Name = "UserId",
                    Type = "S",
                },
                new Aws.DynamoDB.Inputs.TableAttributeArgs
                {
                    Name = "GameTitle",
                    Type = "S",
                },
                new Aws.DynamoDB.Inputs.TableAttributeArgs
                {
                    Name = "TopScore",
                    Type = "N",
                },
            },
            BillingMode = "PROVISIONED",
            GlobalSecondaryIndexes = 
            {
                new Aws.DynamoDB.Inputs.TableGlobalSecondaryIndexArgs
                {
                    HashKey = "GameTitle",
                    Name = "GameTitleIndex",
                    NonKeyAttributes = 
                    {
                        "UserId",
                    },
                    ProjectionType = "INCLUDE",
                    RangeKey = "TopScore",
                    ReadCapacity = 10,
                    WriteCapacity = 10,
                },
            },
            HashKey = "UserId",
            RangeKey = "GameTitle",
            ReadCapacity = 20,
            Tags = 
            {
                { "Environment", "production" },
                { "Name", "dynamodb-table-1" },
            },
            Ttl = new Aws.DynamoDB.Inputs.TableTtlArgs
            {
                AttributeName = "TimeToExist",
                Enabled = false,
            },
            WriteCapacity = 20,
        });
    }

}

