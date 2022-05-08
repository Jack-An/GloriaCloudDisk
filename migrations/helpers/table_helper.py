from alembic import op
from sqlalchemy import text


def table_exists(table_name):
    conn = op.get_bind()
    sql = (
        "SELECT table_name FROM INFORMATION_SCHEMA.tables "
        "WHERE table_schema = DATABASE() AND table_name = '%s';" % table_name
    )
    result = conn.execute(sql)
    return result.returns_rows and result.first() is not None


def tables_exists(table_names):
    conn = op.get_bind()
    sql = (
        "SELECT table_name FROM INFORMATION_SCHEMA.tables "
        "WHERE table_schema = DATABASE() AND table_name in :table_names order by table_name;"
    )
    query = text(sql).bindparams(table_names=tuple(table_names))
    result = conn.execute(query)
    return result.returns_rows and result.rowcount == len(table_names)


def tables():
    conn = op.get_bind()
    sql = (
        "SELECT table_name FROM INFORMATION_SCHEMA.tables "
        "WHERE table_schema = DATABASE() and table_type = 'BASE TABLE' order by table_name;"
    )
    results = conn.execute(text(sql)).fetchall()
    return [r[0] for r in results]


def column_exists(table_name, column_name):
    conn = op.get_bind()
    sql = (
        "SELECT column_name FROM INFORMATION_SCHEMA.columns "
        "WHERE table_schema = DATABASE() AND table_name = '%s' AND column_name = '%s';"
        % (table_name, column_name)
    )
    result = conn.execute(sql)
    return result.returns_rows and result.first() is not None


def column_like(table_name, suffix=None, prefix=None):
    conn = op.get_bind()
    if suffix:
        column_name = "%{}".format(suffix)
    elif prefix:
        column_name = "{}%".format(prefix)

    sql = (
        "SELECT column_name, is_nullable FROM INFORMATION_SCHEMA.columns "
        "WHERE table_schema = DATABASE() AND table_name = '%s' AND column_name like '%s';"
        % (table_name, column_name)
    )
    return conn.execute(text(sql)).fetchall()


def column_of_specific_type(table_name, data_type):
    conn = op.get_bind()
    sql = (
        "SELECT column_name, is_nullable, column_default FROM INFORMATION_SCHEMA.columns "
        "WHERE table_schema = DATABASE() AND table_name = '%s' AND data_type = '%s';"
        % (table_name, data_type)
    )
    return conn.execute(text(sql)).fetchall()


def toggle_exists(toggle_name):
    conn = op.get_bind()
    sql = (
        "SELECT * FROM feature_toggle " "WHERE feature_toggle_name = '%s'" % toggle_name
    )
    result = conn.execute(sql)
    return result.returns_rows and result.first() is not None


def get_column_info(table_name, column_name):
    conn = op.get_bind()
    sql = (
        "SELECT * FROM INFORMATION_SCHEMA.columns "
        "WHERE table_schema = DATABASE() AND table_name = '%s' AND column_name = '%s';"
        % (table_name, column_name)
    )
    result = conn.execute(sql)
    if result.returns_rows:
        return result.first()
    return None


def update_column_type(
    table_name: str, column_name: str, column_type: str, is_null: bool
):
    conn = op.get_bind()
    sql = f"ALTER TABLE {table_name} MODIFY {column_name} {column_type} {'null' if is_null else 'not null'};"
    result = conn.execute(sql)
    if result.returns_rows:
        return result.first()
    return None


def constraint_exists(table_name, constraint_name):
    conn = op.get_bind()
    sql = (
        "SELECT * FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS where TABLE_NAME='%s' and CONSTRAINT_NAME='%s' and CONSTRAINT_SCHEMA=DATABASE();"
        % (table_name, constraint_name)
    )
    result = conn.execute(sql)
    return result.returns_rows and result.first() is not None
