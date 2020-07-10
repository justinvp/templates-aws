using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.WafRegional.WebAcl("example", new Aws.WafRegional.WebAclArgs
        {
            LoggingConfiguration = new Aws.WafRegional.Inputs.WebAclLoggingConfigurationArgs
            {
                LogDestination = aws_kinesis_firehose_delivery_stream.Example.Arn,
                RedactedFields = new Aws.WafRegional.Inputs.WebAclLoggingConfigurationRedactedFieldsArgs
                {
                    FieldToMatch = 
                    {
                        
                        {
                            { "type", "URI" },
                        },
                        
                        {
                            { "data", "referer" },
                            { "type", "HEADER" },
                        },
                    },
                },
            },
        });
    }

}

