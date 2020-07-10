using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.Trigger("example", new Aws.Glue.TriggerArgs
        {
            Actions = 
            {
                new Aws.Glue.Inputs.TriggerActionArgs
                {
                    JobName = aws_glue_job.Example.Name,
                },
            },
            Type = "ON_DEMAND",
        });
    }

}

