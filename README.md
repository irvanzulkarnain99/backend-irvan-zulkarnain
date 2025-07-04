# backend-irvan-zulkarnain

# clone repository

# create db --> marketplace

# create table merchants 

<!-- CREATE TABLE public.merchants (
	id serial4 NOT NULL,
	"name" varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	"password" text NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT merchants_email_key UNIQUE (email),
	CONSTRAINT merchants_pkey PRIMARY KEY (id)
); -->

# create table orders

<!-- CREATE TABLE public.orders (
	id serial4 NOT NULL,
	product_id int4 NOT NULL,
	merchant_id int4 NULL,
	quantity int4 NOT NULL,
	total_price int4 NOT NULL DEFAULT 0,
	shipping_price int4 NOT NULL DEFAULT 0,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	discount int4 NOT NULL DEFAULT 0,
	CONSTRAINT orders_pkey PRIMARY KEY (id)
);
-- public.orders foreign keys
ALTER TABLE public.orders ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE; -->

# create table products

<!-- CREATE TABLE public.products (
	id serial4 NOT NULL,
	merchant_id int4 NULL,
	"name" varchar(100) NOT NULL,
	price int4 NOT NULL,
	stock int4 NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT products_pkey PRIMARY KEY (id)
);
-- public.products foreign keys
ALTER TABLE public.products ADD CONSTRAINT products_merchant_id_fkey FOREIGN KEY (merchant_id) REFERENCES public.merchants(id) ON DELETE CASCADE; -->

<!-- untuk menjalankan project  -->
# go mod tidy
# go run main.go