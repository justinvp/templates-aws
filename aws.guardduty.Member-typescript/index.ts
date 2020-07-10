import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const master = new aws.guardduty.Detector("master", {enable: true});
const memberDetector = new aws.guardduty.Detector("memberDetector", {enable: true}, {
    provider: "aws.dev",
});
const memberMember = new aws.guardduty.Member("memberMember", {
    accountId: memberDetector.accountId,
    detectorId: master.id,
    email: "required@example.com",
    invite: true,
    invitationMessage: "please accept guardduty invitation",
});

