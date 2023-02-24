# query 1
df.lazy().filter((pl.col("integer") > 1) & (pl.col("float") < 2.5))
# query 2
df.lazy().filter(pl.col("integer") > 1).filter(pl.col("float") < 2.5)
