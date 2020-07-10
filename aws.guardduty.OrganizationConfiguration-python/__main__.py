import pulumi
import pulumi_aws as aws

example_detector = aws.guardduty.Detector("exampleDetector", enable=True)
example_organization_configuration = aws.guardduty.OrganizationConfiguration("exampleOrganizationConfiguration",
    auto_enable=True,
    detector_id=example_detector.id)

