using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleTable = new Aws.DynamoDB.Table("exampleTable", new Aws.DynamoDB.TableArgs
        {
            Attributes = 
            {
                new Aws.DynamoDB.Inputs.TableAttributeArgs
                {
                    Name = "exampleHashKey",
                    Type = "S",
                },
            },
            HashKey = "exampleHashKey",
            ReadCapacity = 10,
            WriteCapacity = 10,
        });
        var exampleTableItem = new Aws.DynamoDB.TableItem("exampleTableItem", new Aws.DynamoDB.TableItemArgs
        {
            HashKey = exampleTable.HashKey,
            Item = @"{
  ""exampleHashKey"": {""S"": ""something""},
  ""one"": {""N"": ""11111""},
  ""two"": {""N"": ""22222""},
  ""three"": {""N"": ""33333""},
  ""four"": {""N"": ""44444""}
}

",
            TableName = exampleTable.Name,
        });
    }

}

