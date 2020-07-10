using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.Model("example", new Aws.ApiGatewayV2.ModelArgs
        {
            ApiId = aws_apigatewayv2_api.Example.Id,
            ContentType = "application/json",
            Schema = @"{
  ""$schema"": ""http://json-schema.org/draft-04/schema#"",
  ""title"": ""ExampleModel"",
  ""type"": ""object"",
  ""properties"": {
    ""id"": { ""type"": ""string"" }
  }
}

",
        });
    }

}

