import os
import pandas as pd

print(f'======= start')

local_dir = 'exportdata'

for root, dirs, files in os.walk(local_dir):
    for file_name in files:
        print(f'filename: {file_name}')
        if file_name.endswith('.parquet'):
            parquet_file = os.path.join(root, file_name)
            csv_file = os.path.join(root, file_name.replace('.parquet', '.csv'))

            df = pd.read_parquet(parquet_file)

            df.to_csv(csv_file, index=False)
            print(f'Converted {parquet_file} to {csv_file}')
