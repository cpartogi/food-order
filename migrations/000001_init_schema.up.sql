CREATE TABLE menu_types (
	id uuid NOT NULL,
	menu_type_name varchar(50) NOT NULL,
    created_by varchar(255) NOT NULL,
	created_at timestamptz NOT NULL,
	updated_by varchar(255) NULL,
	updated_at timestamptz NULL,
	CONSTRAINT menu_types PRIMARY KEY (id)
);

CREATE TABLE menus (
	id uuid NOT NULL,
    menu_type_id uuid NOT NULL,
    menu_name varchar(255) NOT NULL,
    menu_description varchar(500) DEFAULT NULL,
    menu_price int4 NULL,
    created_by varchar(255) NOT NULL,
	created_at timestamptz NOT NULL,
	updated_by varchar(255) NULL,
	updated_at timestamptz NULL,
    CONSTRAINT menus PRIMARY KEY (id)
);

create table orders (
	id uuid NOT NULL,
    order_number varchar(50),
    table_number int4 NOT NULL, 
    order_item_id uuid NOT NULL,
    created_by varchar(255) NOT NULL,
	created_at timestamptz NOT NULL,
    CONSTRAINT orders PRIMARY KEY (id)
);

create table order_items (
	id uuid NOT NULL,
    menu_id uuid NOT NULL,
    amount int4 NOT NULL,
    created_by varchar(255) NOT NULL,
	created_at timestamptz NOT NULL,
    CONSTRAINT order_items PRIMARY KEY (id)
);