import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mainIdentityPool = new aws.cognito.IdentityPool("main", {
    allowUnauthenticatedIdentities: false,
    identityPoolName: "identity pool",
    supportedLoginProviders: {
        "graph.facebook.com": "7346241598935555",
    },
});
const authenticatedRole = new aws.iam.Role("authenticated", {
    assumeRolePolicy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "cognito-identity.amazonaws.com"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "cognito-identity.amazonaws.com:aud": "${mainIdentityPool.id}"
        },
        "ForAnyValue:StringLike": {
          "cognito-identity.amazonaws.com:amr": "authenticated"
        }
      }
    }
  ]
}
`,
});
const authenticatedRolePolicy = new aws.iam.RolePolicy("authenticated", {
    policy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "mobileanalytics:PutEvents",
        "cognito-sync:*",
        "cognito-identity:*"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}
`,
    role: authenticatedRole.id,
});
const mainIdentityPoolRoleAttachment = new aws.cognito.IdentityPoolRoleAttachment("main", {
    identityPoolId: mainIdentityPool.id,
    roleMappings: [{
        ambiguousRoleResolution: "AuthenticatedRole",
        identityProvider: "graph.facebook.com",
        mappingRules: [{
            claim: "isAdmin",
            matchType: "Equals",
            roleArn: authenticatedRole.arn,
            value: "paid",
        }],
        type: "Rules",
    }],
    roles: {
        authenticated: authenticatedRole.arn,
    },
});

