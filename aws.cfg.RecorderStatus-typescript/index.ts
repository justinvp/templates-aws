import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.Bucket("b", {});
const fooDeliveryChannel = new aws.cfg.DeliveryChannel("foo", {
    s3BucketName: bucket.bucket,
});
const role = new aws.iam.Role("r", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "config.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
`,
});
const fooRecorder = new aws.cfg.Recorder("foo", {
    roleArn: role.arn,
});
const fooRecorderStatus = new aws.cfg.RecorderStatus("foo", {
    isEnabled: true,
}, { dependsOn: [fooDeliveryChannel] });
const rolePolicyAttachment = new aws.iam.RolePolicyAttachment("a", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSConfigRole",
    role: role.name,
});
const rolePolicy = new aws.iam.RolePolicy("p", {
    policy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:*"
      ],
      "Effect": "Allow",
      "Resource": [
        "${bucket.arn}",
        "${bucket.arn}/*"
      ]
    }
  ]
}
`,
    role: role.id,
});

