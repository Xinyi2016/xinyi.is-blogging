df.select(
    pl.col("string").str.strptime(pl.Date, fmt="%Y-%m-%d")
)
