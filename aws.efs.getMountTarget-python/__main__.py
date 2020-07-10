import pulumi
import pulumi_aws as aws

config = pulumi.Config()
mount_target_id = config.get("mountTargetId")
if mount_target_id is None:
    mount_target_id = ""
by_id = aws.efs.get_mount_target(mount_target_id=mount_target_id)

