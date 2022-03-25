drop table ccca.stock_entry;
drop table ccca.tax_table;

alter table ccca.order
drop
column taxes;
