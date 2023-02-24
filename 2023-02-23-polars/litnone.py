df.with_columns(pl.lit(None).cast(pl.Int64).alias("nulls"))
