using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var www = new Aws.Route53.Record("www", new Aws.Route53.RecordArgs
        {
            Name = "www.example.com",
            Records = 
            {
                aws_eip.Lb.Public_ip,
            },
            Ttl = 300,
            Type = "A",
            ZoneId = aws_route53_zone.Primary.Zone_id,
        });
    }

}

