import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const central = new aws.Provider("central", {
    region: "eu-central-1",
});
const replicationRole = new aws.iam.Role("replication", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "s3.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
`,
});
const destination = new aws.s3.Bucket("destination", {
    region: "eu-west-1",
    versioning: {
        enabled: true,
    },
});
const bucket = new aws.s3.Bucket("bucket", {
    acl: "private",
    region: "eu-central-1",
    replicationConfiguration: {
        role: replicationRole.arn,
        rules: [{
            destination: {
                bucket: destination.arn,
                storageClass: "STANDARD",
            },
            id: "foobar",
            prefix: "foo",
            status: "Enabled",
        }],
    },
    versioning: {
        enabled: true,
    },
}, { provider: central });
const replicationPolicy = new aws.iam.Policy("replication", {
    policy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:GetReplicationConfiguration",
        "s3:ListBucket"
      ],
      "Effect": "Allow",
      "Resource": [
        "${bucket.arn}"
      ]
    },
    {
      "Action": [
        "s3:GetObjectVersion",
        "s3:GetObjectVersionAcl"
      ],
      "Effect": "Allow",
      "Resource": [
        "${bucket.arn}/*"
      ]
    },
    {
      "Action": [
        "s3:ReplicateObject",
        "s3:ReplicateDelete"
      ],
      "Effect": "Allow",
      "Resource": "${destination.arn}/*"
    }
  ]
}
`,
});
const replicationRolePolicyAttachment = new aws.iam.RolePolicyAttachment("replication", {
    policyArn: replicationPolicy.arn,
    role: replicationRole.name,
});

