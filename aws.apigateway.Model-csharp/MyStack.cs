using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myDemoAPI = new Aws.ApiGateway.RestApi("myDemoAPI", new Aws.ApiGateway.RestApiArgs
        {
            Description = "This is my API for demonstration purposes",
        });
        var myDemoModel = new Aws.ApiGateway.Model("myDemoModel", new Aws.ApiGateway.ModelArgs
        {
            ContentType = "application/json",
            Description = "a JSON schema",
            RestApi = myDemoAPI.Id,
            Schema = @"{
  ""type"": ""object""
}

",
        });
    }

}

