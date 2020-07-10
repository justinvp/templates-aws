using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DataSync.NfsLocation("example", new Aws.DataSync.NfsLocationArgs
        {
            OnPremConfig = new Aws.DataSync.Inputs.NfsLocationOnPremConfigArgs
            {
                AgentArns = 
                {
                    aws_datasync_agent.Example.Arn,
                },
            },
            ServerHostname = "nfs.example.com",
            Subdirectory = "/exported/path",
        });
    }

}

