df.with_columns(
    pl.col("qty").rank("dense").alias("rank"),
    pl.col("qty").rank("dense").over("sellerid").alias("group_rank"),
)
