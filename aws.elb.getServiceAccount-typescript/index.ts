import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = pulumi.output(aws.elb.getServiceAccount({ async: true }));
const elbLogs = new aws.s3.Bucket("elb_logs", {
    acl: "private",
    policy: pulumi.interpolate`{
  "Id": "Policy",
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:PutObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::my-elb-tf-test-bucket/AWSLogs/*",
      "Principal": {
        "AWS": [
          "${main.arn}"
        ]
      }
    }
  ]
}
`,
});
const bar = new aws.elb.LoadBalancer("bar", {
    accessLogs: {
        bucket: elbLogs.bucket,
        interval: 5,
    },
    availabilityZones: ["us-west-2a"],
    listeners: [{
        instancePort: 8000,
        instanceProtocol: "http",
        lbPort: 80,
        lbProtocol: "http",
    }],
});

