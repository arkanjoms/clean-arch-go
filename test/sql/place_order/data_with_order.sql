insert into ccca.item (id, category, description, price, width, height, length, weight)
values ('5549d46f-20d3-4d48-9cbe-80acc2b5cbb9', 'Instrumentos Musicais', 'Guitarra', 1000, 100, 50, 15, 3);
insert into ccca.item (id, category, description, price, width, height, length, weight)
values ('cf3dfb32-f654-42b6-be0b-d698eae8a146', 'Instrumentos Musicais', 'Amplificador', 5000, 50, 50, 50, 22);
insert into ccca.item (id, category, description, price, width, height, length, weight)
values ('36ed8660-feaa-4add-94c5-441792e8a0c2', 'Acessórios', 'Cabo', 30, 10, 10, 10, 1);

insert into ccca.order (id, code, cpf, issue_date, shipping_cost, serial)
values ('94c1840a-e8fb-415b-a61b-5bcacb82be3a', '202100000009', '08963482707', '2021-08-20', '10', 9);
insert into ccca.order_item(id_order, id_item, price, quantity)
values ('94c1840a-e8fb-415b-a61b-5bcacb82be3a', '36ed8660-feaa-4add-94c5-441792e8a0c2', 30, 1);
