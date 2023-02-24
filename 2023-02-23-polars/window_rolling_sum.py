df.with_columns(
    pl.col("qty").rolling_sum(window_size=2).alias("rolling_sum"),
    pl.col("qty").rolling_sum(window_size=2).over(
        "sellerid").alias("group_rolling_sum"),
)
