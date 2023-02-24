df.lazy().filter(pl.col("integer") > 1).collect()
