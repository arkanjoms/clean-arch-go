-- COUPON
insert into ccca.coupon (code, percentage, expire_date)
values ('VALE20', 20.0, now() + INTERVAL '1' month);
insert into ccca.coupon (code, percentage, expire_date)
values ('VALE20_EXPIRED', 20.0, now() - INTERVAL '1' month);
