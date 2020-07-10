import pulumi
import pulumi_aws as aws

example = aws.backup.get_selection(plan_id=data["aws_backup_plan"]["example"]["id"],
    selection_id="selection-id-example")

