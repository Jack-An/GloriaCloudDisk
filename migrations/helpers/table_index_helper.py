from alembic import op, context
import sys


def add_index(index_name, table_name, column_list, unique=False):
    if get_index_info(table_name, index_name) is None:
        table_row_count = get_table_row_count(table_name)
        threshold = int(
            context.get_x_argument(as_dictionary=True).get(
                "skip_add_index_row_threshold", 10000000
            )
        )
        if table_row_count > threshold:
            sql = _generate_sql_for_add_index(
                index_name, table_name, column_list, unique
            )
            print_sql(
                "-- Table {} row count {} exceed {}\n{}".format(
                    table_name, table_row_count, threshold, sql
                )
            )
        else:
            print("creating index '%s' to table '%s'" % (index_name, table_name))
            _do_add_index(index_name, table_name, column_list, unique)


def _do_add_index(index_name, table_name, column_list, unique=False):
    if unique:
        op.create_unique_constraint(index_name, table_name, column_list)
    else:
        op.create_index(index_name, table_name, column_list)


def _generate_sql_for_add_index(index_name, table_name, column_list, unique=False):
    index_key_parts = ", ".join(column_list)
    if unique:
        sql = "ALTER TABLE {} ADD UNIQUE KEY {}({});".format(
            table_name, index_name, index_key_parts
        )
    else:
        sql = "CREATE INDEX {} ON {}({});".format(
            index_name, table_name, index_key_parts
        )
    return sql


def rename_index(old_index_name, new_index_name, table_name):
    conn = op.get_bind()
    sql = "ALTER TABLE {} RENAME INDEX {} TO {};".format(
        table_name, old_index_name, new_index_name
    )
    result = conn.execute(sql)
    if result.returns_rows:
        return result.first()
    return None


def remove_index(index_name, table_name):
    if get_index_info(table_name, index_name) is not None:
        print("dropping index '%s' from table '%s'" % (index_name, table_name))
        op.drop_index(index_name, table_name)


def get_index_info(table_name, index_name):
    conn = op.get_bind()
    sql = "show index from `%s` where Key_name = '%s'" % (table_name, index_name)
    result = conn.execute(sql)
    if result.returns_rows:
        return result.first()
    return None


def index_exists(table_name, index_name):
    return get_index_info(table_name, index_name) is not None


def get_table_row_count(table_name):
    conn = op.get_bind()
    sql = (
        "SELECT table_rows FROM INFORMATION_SCHEMA.tables "
        "WHERE table_schema = DATABASE() AND table_name = '%s';" % table_name
    )
    result = conn.execute(sql)
    row_count = result.scalar()
    return -1 if row_count is None else row_count


def print_sql(sql):
    print(
        """
============================ execute sql manually ============================
{}
============================ EXECUTE SQL MANUALLY ============================
""".format(
            sql
        ),
        file=sys.stderr,
    )
