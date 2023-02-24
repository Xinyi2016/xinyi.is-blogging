import polars as pl
import numpy as np
import datetime as dt

today = dt.datetime.utcnow()

df = pl.DataFrame({"integer": np.arange(8),
                   "float": np.arange(0, 4, 0.5, dtype=float),
                   "string": [(today-dt.timedelta(days=i)).strftime("%Y-%m-%d") for i in range(8)],
                   })
