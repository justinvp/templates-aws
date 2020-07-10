using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DataSync.LocationSmb("example", new Aws.DataSync.LocationSmbArgs
        {
            AgentArns = 
            {
                aws_datasync_agent.Example.Arn,
            },
            Password = "ANotGreatPassword",
            ServerHostname = "smb.example.com",
            Subdirectory = "/exported/path",
            User = "Guest",
        });
    }

}

