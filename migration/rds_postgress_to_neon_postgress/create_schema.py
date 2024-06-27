import json

json_file_path = 'exportdata/export_tables_info_folio-db-export_from_1_to_6.json'

with open(json_file_path, 'r') as file:
    schema_data = json.load(file)

type_mapping = {
    "uuid": "UUID",
    "text": "TEXT",
    "timestamp with time zone": "TIMESMAPTZ",
    "bigint": "BIGINT",
    "boolean": "BOOLEAN",
    "ARRAY": "TEXT[]"
}

sql_statements = []

for table_status in schema_data['perTableStatus']:
    if 'schemaMetadata' in table_status:
        table_name = table_status['target'].split('.')[-1]
        columns = table_status['schemaMetadata']['originalTypeMappings']

        column_definitions = []
        for column in columns:
            column_name = column['columnName']
            original_type = column['originalType']
            mapped_type = type_mapping.get(original_type, 'TEXT')
            column_definitions.append(f"{column_name} {mapped_type}")

        column_definitions_str = ",\n ".join(column_definitions)
        create_table_statement = f"CREATE TABLE {table_name} (\n {column_definitions_str}\n);"
        sql_statements.append(create_table_statement)
    else:
        print(f"'schemaMetadata' key not found for table {table_status['target']}")


with open('create_tables.sql', 'w') as f:
    f.write("\n\n".join(sql_statements))

