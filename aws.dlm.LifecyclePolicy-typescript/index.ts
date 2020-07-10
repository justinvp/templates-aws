import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dlmLifecycleRole = new aws.iam.Role("dlm_lifecycle_role", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "dlm.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
`,
});
const dlmLifecycle = new aws.iam.RolePolicy("dlm_lifecycle", {
    policy: `{
   "Version": "2012-10-17",
   "Statement": [
      {
         "Effect": "Allow",
         "Action": [
            "ec2:CreateSnapshot",
            "ec2:DeleteSnapshot",
            "ec2:DescribeVolumes",
            "ec2:DescribeSnapshots"
         ],
         "Resource": "*"
      },
      {
         "Effect": "Allow",
         "Action": [
            "ec2:CreateTags"
         ],
         "Resource": "arn:aws:ec2:*::snapshot/*"
      }
   ]
}
`,
    role: dlmLifecycleRole.id,
});
const example = new aws.dlm.LifecyclePolicy("example", {
    description: "example DLM lifecycle policy",
    executionRoleArn: dlmLifecycleRole.arn,
    policyDetails: {
        resourceTypes: ["VOLUME"],
        schedules: [{
            copyTags: false,
            createRule: {
                interval: 24,
                intervalUnit: "HOURS",
                times: "23:45",
            },
            name: "2 weeks of daily snapshots",
            retainRule: {
                count: 14,
            },
            tagsToAdd: {
                SnapshotCreator: "DLM",
            },
        }],
        targetTags: {
            Snapshot: "true",
        },
    },
    state: "ENABLED",
});

