pl.DataFrame({"a": [1, 2, 3], "b": [4, 5, 6]}).select(
    [
        pl.all().log().suffix("_polars_log"),
        np.log(pl.all()).suffix("_np_log"),
    ]
)
