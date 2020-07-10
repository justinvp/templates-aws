using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleRole = new Aws.Iam.Role("exampleRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": [""sts:AssumeRole""],
      ""Effect"": ""allow"",
      ""Principal"": {
        ""Service"": [""backup.amazonaws.com""]
      }
    }
  ]
}

",
        });
        var exampleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("exampleRolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup",
            Role = exampleRole.Name,
        });
        var exampleSelection = new Aws.Backup.Selection("exampleSelection", new Aws.Backup.SelectionArgs
        {
            IamRoleArn = exampleRole.Arn,
        });
    }

}

