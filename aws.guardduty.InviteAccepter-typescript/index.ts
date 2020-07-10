import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const master = new aws.guardduty.Detector("master", {});
const memberDetector = new aws.guardduty.Detector("memberDetector", {}, {
    provider: "aws.dev",
});
const dev = new aws.guardduty.Member("dev", {
    accountId: memberDetector.accountId,
    detectorId: master.id,
    email: "required@example.com",
    invite: true,
});
const memberInviteAccepter = new aws.guardduty.InviteAccepter("memberInviteAccepter", {
    detectorId: memberDetector.id,
    masterAccountId: master.accountId,
}, {
    provider: "aws.dev",
    dependsOn: ["aws_guardduty_member.dev"],
});

