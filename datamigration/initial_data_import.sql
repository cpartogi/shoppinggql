INSERT INTO checkoutpromo.customers (customer_id,customer_name,customer_email) VALUES
	 ('32dfc16b-6df8-4316-b3a8-b50cb3a6b651','john doe','jhondoe@anywhere.com');

INSERT INTO checkoutpromo.orders (order_id,order_num,customer_id,product_id,unit_price,qty,total_price,created_at) VALUES
	 ('4799054c-36d6-46f2-489c-2f9d17df9a23','INV-2022-4-18-11-38-29','32dfc16b-6df8-4316-b3a8-b50cb3a6b651','30acdd4e-ea07-4a0b-b2a0-dbc403617e60',49.99,1,49.99,'2022-04-18 11:38:29.0'),
	 ('6aeb7cdc-06c4-46d9-50a9-f2b52b8ffb29','INV-2022-4-18-14-16-48','32dfc16b-6df8-4316-b3a8-b50cb3a6b651','30acdd4e-ea07-4a0b-b2a0-dbc403617e60',33.3267,3,99.9801,'2022-04-18 14:16:48.0'),
	 ('83391851-9aae-4c94-65b8-bc088b48f772','INV-2022-4-18-14-50-52','32dfc16b-6df8-4316-b3a8-b50cb3a6b651','8a870301-110c-4d2c-99c0-99df2216672c',0.0,1,0.0,'2022-04-18 14:50:52.0'),
	 ('856344c3-66f6-4554-45f4-4bbb097a6905','INV-2022-4-18-14-37-8','32dfc16b-6df8-4316-b3a8-b50cb3a6b651','8a870301-110c-4d2c-99c0-99df2216672c',0.0,1,0.0,'2022-04-18 14:37:08.0'),
	 ('b00164b5-9b57-4c55-4310-a13e98690bae','INV-2022-4-18-14-37-8','32dfc16b-6df8-4316-b3a8-b50cb3a6b651','bcedbeb4-ef91-4452-adeb-421de1ec1058',5399.99,1,5399.99,'2022-04-18 14:37:08.0'),
	 ('b618b96f-4908-4ee8-6b36-93efcd8852ce','INV-2022-4-18-12-7-57','32dfc16b-6df8-4316-b3a8-b50cb3a6b651','30acdd4e-ea07-4a0b-b2a0-dbc403617e60',49.99,1,49.99,'2022-04-18 12:07:57.0'),
	 ('bb935709-5964-44b3-58ef-361cadd38724','INV-2022-4-18-12-7-57','32dfc16b-6df8-4316-b3a8-b50cb3a6b651','8a870301-110c-4d2c-99c0-99df2216672c',30.0,1,30.0,'2022-04-18 12:07:57.0'),
	 ('c977a142-aa45-4772-7691-c3b056cfcba3','INV-2022-4-18-14-50-52','32dfc16b-6df8-4316-b3a8-b50cb3a6b651','bcedbeb4-ef91-4452-adeb-421de1ec1058',5399.99,1,5399.99,'2022-04-18 14:50:52.0'),
	 ('debac501-7782-42dc-670c-eaff44ab93fe','INV-2022-4-18-11-38-29','32dfc16b-6df8-4316-b3a8-b50cb3a6b651','8a870301-110c-4d2c-99c0-99df2216672c',30.0,1,30.0,'2022-04-18 11:38:29.0'),
	 ('fffb9ffe-73d7-401e-7fa6-d99e282b7e0e','INV-2022-4-18-14-12-7','32dfc16b-6df8-4316-b3a8-b50cb3a6b651','0a2f3e0d-6d80-43f2-8adc-827ce52f2510',98.55,3,295.65,'2022-04-18 14:12:07.0'); 

INSERT INTO checkoutpromo.products (product_id,product_sku,product_name,product_price,product_qty) VALUES
	 ('0a2f3e0d-6d80-43f2-8adc-827ce52f2510','A304SD','Alexa Speaker',109.5,10),
	 ('30acdd4e-ea07-4a0b-b2a0-dbc403617e60','120P90','Google Home',49.99,10),
	 ('8a870301-110c-4d2c-99c0-99df2216672c','234234','Raspberry Pi B',30.0,2),
	 ('bcedbeb4-ef91-4452-adeb-421de1ec1058','43N23P','MacBook Pro',5399.99,5);

INSERT INTO checkoutpromo.promo_rules (promo_id,promo_name,product_id,min_qty,price,bonus_product_id) VALUES
	 ('177240a0-1523-42a9-8d41-6562c18ea495','Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers','0a2f3e0d-6d80-43f2-8adc-827ce52f2510',3,98.55,NULL),
	 ('32714d2a-18dd-4826-bf0e-374f7bd5de13','Buy 3 Google Homes for the price of 2','30acdd4e-ea07-4a0b-b2a0-dbc403617e60',3,33.3267,NULL),
	 ('95dbff02-602b-4754-a2c1-f06a74753ea8','Each sale of a MacBook Pro comes with a free Raspberry Pi B','bcedbeb4-ef91-4452-adeb-421de1ec1058',1,0.0,'8a870301-110c-4d2c-99c0-99df2216672c');     
 
     