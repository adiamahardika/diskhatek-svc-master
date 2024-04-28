-- DROP SCHEMA public;

CREATE SCHEMA public AUTHORIZATION pg_database_owner;

COMMENT ON SCHEMA public IS 'standard public schema';

-- DROP TYPE public.order_types;

CREATE TYPE public.order_types AS ENUM (
	'pending',
	'processing',
	'completed',
	'cancelled');

-- DROP TYPE public.warehouse_status;

CREATE TYPE public.warehouse_status AS ENUM (
	'active',
	'inactive');

-- DROP SEQUENCE public.order_item_order_item_id_seq;

CREATE SEQUENCE public.order_item_order_item_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.orders_order_id_seq;

CREATE SEQUENCE public.orders_order_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.products_product_id_seq;

CREATE SEQUENCE public.products_product_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.reserved_stock_reservation_id_seq;

CREATE SEQUENCE public.reserved_stock_reservation_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.shops_shop_id_seq;

CREATE SEQUENCE public.shops_shop_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.stock_id_seq;

CREATE SEQUENCE public.stock_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.stock_transfer_transfer_id_seq;

CREATE SEQUENCE public.stock_transfer_transfer_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.users_user_id_seq;

CREATE SEQUENCE public.users_user_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.warehouse_warehouse_id_seq;

CREATE SEQUENCE public.warehouse_warehouse_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;-- public.shops definition

-- Drop table

-- DROP TABLE public.shops;

CREATE TABLE public.shops (
	shop_id serial4 NOT NULL,
	"name" varchar(255) NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT shops_pkey PRIMARY KEY (shop_id)
);


-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	user_id serial4 NOT NULL,
	"name" varchar(255) NOT NULL,
	phone varchar(20) NOT NULL,
	password_hash varchar(255) NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	email varchar(20) NOT NULL,
	CONSTRAINT users_email_un UNIQUE (email),
	CONSTRAINT users_phone_un UNIQUE (phone),
	CONSTRAINT users_pkey PRIMARY KEY (user_id)
);


-- public.orders definition

-- Drop table

-- DROP TABLE public.orders;

CREATE TABLE public.orders (
	order_id serial4 NOT NULL,
	user_id int4 NOT NULL,
	order_date date NOT NULL DEFAULT CURRENT_TIMESTAMP,
	total_amount numeric(10, 2) NOT NULL,
	status public.order_types NOT NULL DEFAULT 'pending'::order_types,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	payment_deadline timestamp NOT NULL,
	CONSTRAINT orders_pkey PRIMARY KEY (order_id),
	CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id)
);


-- public.products definition

-- Drop table

-- DROP TABLE public.products;

CREATE TABLE public.products (
	product_id serial4 NOT NULL,
	"name" varchar(255) NULL,
	description text NULL,
	price numeric(10, 2) NULL,
	shop_id int4 NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT products_pkey PRIMARY KEY (product_id),
	CONSTRAINT products_shop_id_fkey FOREIGN KEY (shop_id) REFERENCES public.shops(shop_id)
);


-- public.warehouses definition

-- Drop table

-- DROP TABLE public.warehouses;

CREATE TABLE public.warehouses (
	warehouse_id int4 NOT NULL DEFAULT nextval('warehouse_warehouse_id_seq'::regclass),
	"name" varchar(255) NULL,
	shop_id int4 NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	status public.warehouse_status NOT NULL DEFAULT 'active'::warehouse_status,
	CONSTRAINT warehouse_pkey PRIMARY KEY (warehouse_id),
	CONSTRAINT warehouse_shop_id_fkey FOREIGN KEY (shop_id) REFERENCES public.shops(shop_id)
);


-- public.order_items definition

-- Drop table

-- DROP TABLE public.order_items;

CREATE TABLE public.order_items (
	order_item_id int4 NOT NULL DEFAULT nextval('order_item_order_item_id_seq'::regclass),
	order_id int4 NULL,
	product_id int4 NULL,
	quantity int4 NULL,
	unit_price numeric(10, 2) NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT order_item_pkey PRIMARY KEY (order_item_id),
	CONSTRAINT order_item_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(order_id),
	CONSTRAINT order_item_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(product_id)
);


-- public.reserved_stocks definition

-- Drop table

-- DROP TABLE public.reserved_stocks;

CREATE TABLE public.reserved_stocks (
	reservation_id int4 NOT NULL DEFAULT nextval('reserved_stock_reservation_id_seq'::regclass),
	order_item_id int4 NULL,
	product_id int4 NULL,
	reserved_quantity int4 NULL,
	reservation_expiry_time timestamp NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT reserved_stock_pkey PRIMARY KEY (reservation_id),
	CONSTRAINT reserved_stock_order_item_id_fkey FOREIGN KEY (order_item_id) REFERENCES public.order_items(order_item_id),
	CONSTRAINT reserved_stock_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(product_id)
);


-- public.stock_transfers definition

-- Drop table

-- DROP TABLE public.stock_transfers;

CREATE TABLE public.stock_transfers (
	transfer_id int4 NOT NULL DEFAULT nextval('stock_transfer_transfer_id_seq'::regclass),
	product_id int4 NULL,
	source_warehouse_id int4 NULL,
	destination_warehouse_id int4 NULL,
	quantity int4 NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	transfer_date date NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT stock_transfer_pkey PRIMARY KEY (transfer_id),
	CONSTRAINT stock_transfer_destination_warehouse_id_fkey FOREIGN KEY (destination_warehouse_id) REFERENCES public.warehouses(warehouse_id),
	CONSTRAINT stock_transfer_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(product_id),
	CONSTRAINT stock_transfer_source_warehouse_id_fkey FOREIGN KEY (source_warehouse_id) REFERENCES public.warehouses(warehouse_id)
);


-- public.stocks definition

-- Drop table

-- DROP TABLE public.stocks;

CREATE TABLE public.stocks (
	stock_id int4 NOT NULL DEFAULT nextval('stock_id_seq'::regclass),
	product_id int4 NOT NULL,
	warehouse_id int4 NOT NULL,
	quantity int4 NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT stock_pkey PRIMARY KEY (stock_id),
	CONSTRAINT stock_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(product_id),
	CONSTRAINT stock_warehouse_id_fkey FOREIGN KEY (warehouse_id) REFERENCES public.warehouses(warehouse_id)
);