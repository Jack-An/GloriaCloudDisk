import sqlalchemy as sa  # type: ignore

import sys

from os.path import abspath, dirname
sys.path.append(dirname(dirname(abspath(__file__))))


def create_at_column():
    return sa.Column(
        "created_at", sa.TIMESTAMP, nullable=False, server_default=sa.func.now()
    )


def update_at_column():
    return sa.Column(
        "updated_at",
        sa.TIMESTAMP,
        nullable=False,
        server_default=sa.text("CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"),
    )


def deleted_column():
    return sa.Column("deleted", sa.Boolean, nullable=False, server_default="0")


def org_column():
    return sa.Column("org_code", sa.String(255), nullable=False)
