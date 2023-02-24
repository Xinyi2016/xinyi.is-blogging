  RECHUNK
    UNION:
    PLAN 0:
      CSV SCAN test1.csv
      PROJECT */3 COLUMNS
    PLAN 1:
      CSV SCAN test2.csv
      PROJECT */3 COLUMNS
    PLAN 2:
      CSV SCAN testdata.csv
      PROJECT */3 COLUMNS
    END UNION