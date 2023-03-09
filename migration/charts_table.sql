create table charts (
  exchange text,
  pair text,
  time_frame text,
  open_time bigint,
  open real,
  high real,
  low real,
  close real,
  volume real,
  primary key (exchange, pair, time_frame, open_time)
);
