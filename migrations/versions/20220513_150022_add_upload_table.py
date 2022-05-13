"""add upload table

Revision ID: 409c52f1398a
Revises: 8d443b58d992
Create Date: 2022-05-13 15:00:22.953934

"""
from alembic import op
import sqlalchemy as sa

import sys

from os.path import abspath, dirname
sys.path.append(dirname(dirname(abspath(__file__))))

from migrations.versions import create_at_column, update_at_column, deleted_column

# revision identifiers, used by Alembic.
from migrations.helpers.table_helper import table_exists

revision = '409c52f1398a'
down_revision = '8d443b58d992'
branch_labels = None
depends_on = None


def upgrade():
    if not table_exists("upload"):
        op.create_table(
            "upload",
            sa.Column("id", sa.BIGINT, primary_key=True),
            sa.Column("parent_id", sa.BIGINT),
            sa.Column("name", sa.String(1024), nullable=False),
            sa.Column("type", sa.String(36), nullable=False),
            sa.Column("ext", sa.String(16), nullable=True),
            sa.Column("key", sa.String(512), nullable=True),
            sa.Column("size", sa.INT, nullable=False, server_default="0"),
            create_at_column(),
            update_at_column(),
            deleted_column(),
        )

    if not table_exists("resource_relation"):
        op.create_table("resource_relation",
                        sa.Column("id", sa.BIGINT, primary_key=True),
                        sa.Column("user_id", sa.BIGINT, nullable=False),
                        sa.Column("upload_id", sa.BIGINT, nullable=False),
                        create_at_column(),
                        update_at_column(),
                        deleted_column())


def downgrade():
    pass
