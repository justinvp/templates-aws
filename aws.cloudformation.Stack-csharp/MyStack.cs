using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var network = new Aws.CloudFormation.Stack("network", new Aws.CloudFormation.StackArgs
        {
            Parameters = 
            {
                { "VPCCidr", "10.0.0.0/16" },
            },
            TemplateBody = @"{
  ""Parameters"" : {
    ""VPCCidr"" : {
      ""Type"" : ""String"",
      ""Default"" : ""10.0.0.0/16"",
      ""Description"" : ""Enter the CIDR block for the VPC. Default is 10.0.0.0/16.""
    }
  },
  ""Resources"" : {
    ""myVpc"": {
      ""Type"" : ""AWS::EC2::VPC"",
      ""Properties"" : {
        ""CidrBlock"" : { ""Ref"" : ""VPCCidr"" },
        ""Tags"" : [
          {""Key"": ""Name"", ""Value"": ""Primary_CF_VPC""}
        ]
      }
    }
  }
}

",
        });
    }

}

