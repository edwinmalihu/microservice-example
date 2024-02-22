# microservice-example

Table Customer
CREATE TABLE public.customer (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	username varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	email varchar(255) NULL,
	name varchar(255) NULL,
	CONSTRAINT customer_email_key UNIQUE (email),
	CONSTRAINT customer_pkey PRIMARY KEY (id),
	CONSTRAINT customer_username_key UNIQUE (username)
);
CREATE INDEX idx_customer_deleted_at ON public.customer USING btree (deleted_at);

Table Category
CREATE TABLE public.category (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	category varchar(255) NOT NULL,
	CONSTRAINT category_pkey PRIMARY KEY (id),
	CONSTRAINT uni_category_category UNIQUE (category)
);
CREATE INDEX idx_category_deleted_at ON public.category USING btree (deleted_at);

Table Product
CREATE TABLE public.product (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" varchar(255) NOT NULL,
	description varchar(255) NULL,
	price numeric(22, 2) NULL,
	stok int8 NULL,
	CONSTRAINT product_pkey PRIMARY KEY (id),
	CONSTRAINT uni_product_name UNIQUE (name)
);
CREATE INDEX idx_product_deleted_at ON public.product USING btree (deleted_at);

Table Relasi Product Category

CREATE TABLE public.product_category (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	product_id int8 NULL,
	category_id int8 NULL,
	CONSTRAINT product_category_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_product_category_deleted_at ON public.product_category USING btree (deleted_at);


ALTER TABLE public.product_category ADD CONSTRAINT fk_product_category_category FOREIGN KEY (category_id) REFERENCES public.category(id);
ALTER TABLE public.product_category ADD CONSTRAINT fk_product_category_product FOREIGN KEY (product_id) REFERENCES public.product(id);

Table Cart
CREATE TABLE public.cart (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	customer_id int8 NULL,
	product_id int8 NULL,
	quantity int8 NULL,
	CONSTRAINT cart_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_cart_deleted_at ON public.cart USING btree (deleted_at);

ALTER TABLE public.cart ADD CONSTRAINT fk_cart_customer FOREIGN KEY (customer_id) REFERENCES public.customer(id);
ALTER TABLE public.cart ADD CONSTRAINT fk_cart_product FOREIGN KEY (product_id) REFERENCES public.product(id);

Table Order
CREATE TABLE public."order" (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	customer_id int8 NULL,
	total_amount numeric(22, 2) NULL,
	order_date timestamptz NULL,
	status text NULL,
	CONSTRAINT order_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_order_deleted_at ON public."order" USING btree (deleted_at);

ALTER TABLE public."order" ADD CONSTRAINT fk_order_customer FOREIGN KEY (customer_id) REFERENCES public.customer(id);

Table Order Item
CREATE TABLE public.order_item (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	order_id int8 NULL,
	product_id int8 NULL,
	quantity int8 NULL,
	CONSTRAINT order_item_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_order_item_deleted_at ON public.order_item USING btree (deleted_at);

ALTER TABLE public.order_item ADD CONSTRAINT fk_order_item_order FOREIGN KEY (order_id) REFERENCES public."order"(id);
ALTER TABLE public.order_item ADD CONSTRAINT fk_order_item_product FOREIGN KEY (product_id) REFERENCES public.product(id);

Table Payment
CREATE TABLE public.payments (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	order_id int8 NULL,
	amount numeric NULL,
	payment_date timestamptz NULL,
	status text NULL,
	CONSTRAINT payments_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_payments_deleted_at ON public.payments USING btree (deleted_at);

ALTER TABLE public.payments ADD CONSTRAINT fk_payments_order FOREIGN KEY (order_id) REFERENCES public."order"(id);