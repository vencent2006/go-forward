tables = []
for i in range(128):
    table_name = f"reward_vmember_user_{i:02x}"
    tables.append(table_name)

sql_parts = []
for table in tables:
    part = f"""
SELECT vuid, uid, status, type, is_show, freeze, create_time, start_time, expire_time, extend
FROM {table}
WHERE create_time < '2025-05-20 23:59:59'"""
    sql_parts.append(part)

full_sql = "\nUNION ALL\n".join(sql_parts)
print(full_sql)
