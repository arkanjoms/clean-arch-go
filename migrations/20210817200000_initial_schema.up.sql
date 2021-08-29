create schema ccca;

create table ccca.item
(
    id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    category    text,
    description text,
    price       numeric,
    width       integer,
    height      integer,
    length      integer,
    weight      integer
);

create table ccca.coupon
(
    code        text,
    percentage  numeric,
    expire_date timestamp,
    primary key (code)
);

create table ccca.order
(
    id            uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    coupon_code   text,
    code          text,
    cpf           text,
    issue_date    timestamp,
    shipping_cost numeric,
    serial        integer
);

create table ccca.order_item
(
    id_order uuid,
    id_item  uuid,
    price    numeric,
    quantity integer,
    primary key (id_order, id_item)
);
