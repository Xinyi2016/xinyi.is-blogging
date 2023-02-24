buf = io.BytesIO()
df.write_parquet(buf, compression='uncompressed')
