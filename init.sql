create table deliveries(
    id 		serial primary key,
    name 	text not null,
    phone 	text not null,
    zip 	text not null,
    city 	text not null,
    address text not null,
    region  text not null,
    email   text not null
);

create table payments(
    transaction 	uuid primary key,
    request_id 		text,
    currency 		text not null,
    provider 		text not null,
    amount 			int not null,
    payment_dt 		int not null,
    bank 			text not null,
    delivery_cost	int not null,
    goods_total 	int not null,
    custom_fee 		int not null
);

create table models(
    order_uid 		    uuid primary key,
    track_number 		text unique not null,
    entry 			    text not null,
    delivery_id 		int references deliveries(id) not null,
    payment_id 		    uuid references payments(transaction) not null,
    locale  			text not null,
    internal_signature 	text not null,
    customer_id 	    text not null,
    delivery_service 	text not null,
    shard_key 			text not null,
    sm_id 				int not null,
    date_created 		timestamp without time zone default NOW(),
    oof_shard 			text not null
);

create table items(
    chrt_id 	 int primary key,
    track_number text unique not null,
    price 		 int not null,
    rid 		 text not null,
    name 		 text not null,
    sale 		 int not null,
    size 		 text not null,
    total_price  int not null,
    nm_id		 int not null,
    brand 		 text not null,
    status 		 int not null,
    models_id 	 uuid references models(order_uid) not null
);

insert into deliveries values (
    1,
    'Test Testov',
    '+9720000000',
    '2639809',
    'Kiryat Mozkin',
    'Ploshad Mira 15',
    'Kraiot',
    'test@gmail.com'
);

insert into payments values(
    'ebfe6316-e5f8-f465-38eb-f79b97b798c5',
    'test',
    'USD',
    'wbpay',
    1817,
    1637907727,
    'alpha',
    1500,
    317,
    1
);

insert into models values (
    'ebfe6316-e5f8-f465-38eb-f79b97b798c6',
    'WBILMTESTTRACK',
    'WBIL',
    1,
    'ebfe6316-e5f8-f465-38eb-f79b97b798c5',
    'en',
    'test',
    'test',
    'meest',
    '9',
    99,
    '2021-11-26T06:22:19Z',
    '1'
);

insert into items values (
    9934930,
    'WBILMTESTTRACK',
    453,
    'ab4219087a764ae0btest',
    'Mascaras',
    30,
    '123',
    317,
    2389212,
    'Vivienne Sabo',
    202,
    'ebfe6316-e5f8-f465-38eb-f79b97b798c6'
);