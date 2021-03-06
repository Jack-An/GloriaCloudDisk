"""add basic table

Revision ID: 8d443b58d992
Revises: 
Create Date: 2022-05-08 12:45:19.586580

"""
from alembic import op
import sqlalchemy as sa

import sys

from os.path import abspath, dirname
sys.path.append(dirname(dirname(abspath(__file__))))

from migrations.versions import create_at_column, update_at_column, deleted_column



from migrations.helpers.table_helper import table_exists

# revision identifiers, used by Alembic.
revision = "8d443b58d992"
down_revision = None
branch_labels = None
depends_on = None



def upgrade():
    if not table_exists("user"):
        op.create_table(
            "user",
            sa.Column("id", sa.BIGINT, primary_key=True),
            sa.Column("name", sa.String(255), nullable=False),
            sa.Column("email", sa.String(512), nullable=True),
            sa.Column("phone", sa.String(36), nullable=True),
            sa.Column("source", sa.String(36), nullable=False),  # 注册方式
            sa.Column("active", sa.Boolean, nullable=False, server_default="1"),
            create_at_column(),
            update_at_column(),
            deleted_column(),
        )

    if not table_exists("identity"):
        op.create_table(
            "identity",
            sa.Column("id", sa.BIGINT, primary_key=True),
            sa.Column("password", sa.String(512), nullable=False),
            sa.Column("lock_time", sa.DateTime, nullable=True),
            create_at_column(),
            update_at_column(),
            deleted_column(),
        )


def downgrade():
    if table_exists("user"):
        op.drop_table("user")

    if table_exists("identity"):
        op.drop_table("identity")
