df.with_columns(
    pl.col("qty").rank("ordinal").alias("rank"),
    pl.col("qty").rank("ordinal").over("sellerid").alias("group_rank"),
)
