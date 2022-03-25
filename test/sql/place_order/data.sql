insert into ccca.item (id, category, description, price, width, height, length, weight)
values ('5549d46f-20d3-4d48-9cbe-80acc2b5cbb9', 'Instrumentos Musicais', 'Guitarra', 1000, 100, 50, 15, 3);
insert into ccca.item (id, category, description, price, width, height, length, weight)
values ('cf3dfb32-f654-42b6-be0b-d698eae8a146', 'Instrumentos Musicais', 'Amplificador', 5000, 50, 50, 50, 22);
insert into ccca.item (id, category, description, price, width, height, length, weight)
values ('36ed8660-feaa-4add-94c5-441792e8a0c2', 'Acessórios', 'Cabo', 30, 10, 10, 10, 1);

-- COUPON
insert into ccca.coupon (code, percentage, expire_date)
values ('VALE20', 20.0, now() + INTERVAL '1' month);
insert into ccca.coupon (code, percentage, expire_date)
values ('VALE20_EXPIRED', 20.0, now() - INTERVAL '1' month);
