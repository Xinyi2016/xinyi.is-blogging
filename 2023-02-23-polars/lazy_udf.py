def add_feature1(df):
    return df.with_columns(pl.col("integer").min().alias("min_value"))


def add_feature2(df):
    return df.with_columns(pl.col("string").n_unique().alias("num_string"))


df.lazy().pipe(add_feature1).pipe(add_feature2).collect()
