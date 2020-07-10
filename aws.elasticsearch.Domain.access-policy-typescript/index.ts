import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const domain = config.get("domain") || "tf-test";

const currentRegion = pulumi.output(aws.getRegion({ async: true }));
const currentCallerIdentity = pulumi.output(aws.getCallerIdentity({ async: true }));
const example = new aws.elasticsearch.Domain("example", {
    accessPolicies: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "es:*",
      "Principal": "*",
      "Effect": "Allow",
      "Resource": "arn:aws:es:${currentRegion.name!}:${currentCallerIdentity.accountId}:domain/${domain}/*",
      "Condition": {
        "IpAddress": {"aws:SourceIp": ["66.193.100.22/32"]}
      }
    }
  ]
}
`,
});

