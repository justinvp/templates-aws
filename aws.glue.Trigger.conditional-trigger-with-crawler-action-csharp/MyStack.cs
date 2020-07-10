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
                    CrawlerName = aws_glue_crawler.Example1.Name,
                },
            },
            Predicate = new Aws.Glue.Inputs.TriggerPredicateArgs
            {
                Conditions = 
                {
                    new Aws.Glue.Inputs.TriggerPredicateConditionArgs
                    {
                        JobName = aws_glue_job.Example2.Name,
                        State = "SUCCEEDED",
                    },
                },
            },
            Type = "CONDITIONAL",
        });
    }

}

