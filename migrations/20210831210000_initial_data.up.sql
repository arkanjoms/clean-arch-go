-- TAX_TABLES
create table ccca.tax_table
(
    id      uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    id_item uuid    not null,
    type    text    not null,
    value   numeric not null
);

insert into ccca.tax_table (id_item, type, value)
values ('5549d46f-20d3-4d48-9cbe-80acc2b5cbb9', 'default', 15);
insert into ccca.tax_table (id_item, type, value)
values ('cf3dfb32-f654-42b6-be0b-d698eae8a146', 'default', 15);
insert into ccca.tax_table (id_item, type, value)
values ('36ed8660-feaa-4add-94c5-441792e8a0c2', 'default', 5);
insert into ccca.tax_table (id_item, type, value)
values ('5549d46f-20d3-4d48-9cbe-80acc2b5cbb9', 'november', 5);
insert into ccca.tax_table (id_item, type, value)
values ('cf3dfb32-f654-42b6-be0b-d698eae8a146', 'november', 5);
insert into ccca.tax_table (id_item, type, value)
values ('36ed8660-feaa-4add-94c5-441792e8a0c2', 'november', 1);

-- ORDER
alter table ccca.order
    add column taxes numeric;

-- STOCK_ENTRY
create table ccca.stock_entry
(
    id        uuid PRIMARY KEY   DEFAULT gen_random_uuid(),
    id_item   uuid      not null,
    operation text      not null,
    quantity  integer   not null,
    date      timestamp not null default now()
);

insert into ccca.stock_entry (id_item, operation, quantity)
values ('5549d46f-20d3-4d48-9cbe-80acc2b5cbb9', 'in', 10);
insert into ccca.stock_entry (id_item, operation, quantity)
values ('cf3dfb32-f654-42b6-be0b-d698eae8a146', 'in', 10);
insert into ccca.stock_entry (id_item, operation, quantity)
values ('36ed8660-feaa-4add-94c5-441792e8a0c2', 'in', 10);
