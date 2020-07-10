using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Waf.WebAcl("example", new Aws.Waf.WebAclArgs
        {
            LoggingConfiguration = new Aws.Waf.Inputs.WebAclLoggingConfigurationArgs
            {
                LogDestination = aws_kinesis_firehose_delivery_stream.Example.Arn,
                RedactedFields = new Aws.Waf.Inputs.WebAclLoggingConfigurationRedactedFieldsArgs
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

