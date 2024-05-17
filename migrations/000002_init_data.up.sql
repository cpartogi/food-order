INSERT INTO menu_types (id, menu_type_name, created_by, created_at, updated_by, updated_at) values ('2815a81b-602b-462b-abcc-160d47c74a40', 'Food', 'admin', now(), 'admin', now());

INSERT INTO menu_types (id, menu_type_name, created_by, created_at, updated_by, updated_at) values ('d1ee752a-bee7-4b79-8c27-8156cffb5cd8', 'Beverages', 'admin', now(), 'admin', now());


INSERT INTO menus (id, menu_type_id, menu_name, menu_description, menu_price, created_by, created_at, updated_by, updated_at) values ('8b81b79b-acfb-4955-8546-0770378734e6', '2815a81b-602b-462b-abcc-160d47c74a40', 'Regular Fried Rice', 'Fried rice with egg and onions', 15000,  'admin', now(), 'admin', now());

INSERT INTO menus (id, menu_type_id, menu_name, menu_description, menu_price, created_by, created_at, updated_by, updated_at) values ('c8fed528-29fc-4485-8dd9-fc0c78d804f8', 'd1ee752a-bee7-4b79-8c27-8156cffb5cd8', 'Iced Tea', 'Tea served with ice block', 3000,  'admin', now(), 'admin', now());