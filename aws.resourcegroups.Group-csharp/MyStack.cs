using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.ResourceGroups.Group("test", new Aws.ResourceGroups.GroupArgs
        {
            ResourceQuery = new Aws.ResourceGroups.Inputs.GroupResourceQueryArgs
            {
                Query = @"{
  ""ResourceTypeFilters"": [
    ""AWS::EC2::Instance""
  ],
  ""TagFilters"": [
    {
      ""Key"": ""Stage"",
      ""Values"": [""Test""]
    }
  ]
}

",
            },
        });
    }

}

