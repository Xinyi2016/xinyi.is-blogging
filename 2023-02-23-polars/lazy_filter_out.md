  FILTER [([(col("integer")) > (1i32)]) & ([(col("float")) < (2.5f64)])] FROM
    DF ["integer", "float", "string"]; PROJECT */3 COLUMNS; SELECTION: "None"