using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Ssm.Document("foo", new Aws.Ssm.DocumentArgs
        {
            Content = @"  {
    ""schemaVersion"": ""1.2"",
    ""description"": ""Check ip configuration of a Linux instance."",
    ""parameters"": {

    },
    ""runtimeConfig"": {
      ""aws:runShellScript"": {
        ""properties"": [
          {
            ""id"": ""0.aws:runShellScript"",
            ""runCommand"": [""ifconfig""]
          }
        ]
      }
    }
  }

",
            DocumentType = "Command",
        });
    }

}

