import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.elasticsearch.Domain("example", {
    elasticsearchVersion: "2.3",
});
const main = new aws.elasticsearch.DomainPolicy("main", {
    accessPolicies: pulumi.interpolate`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "es:*",
            "Principal": "*",
            "Effect": "Allow",
            "Condition": {
                "IpAddress": {"aws:SourceIp": "127.0.0.1/32"}
            },
            "Resource": "${example.arn}/*"
        }
    ]
}
`,
    domainName: example.domainName,
});

