INSERT INTO items (
  id, name, description, price, currency, status, stock, images, metadata
)
VALUES
('itm_001','Deluxe Ocean Room','Luxury room with ocean view and balcony',15000,'USD','active',10,
 ARRAY['https://images.unsplash.com/photo-1501117716987-c8e1ecb210d1'],
 '{"bed":"king","view":"ocean"}'),

('itm_002','Standard Room','Comfortable budget-friendly room',8000,'USD','active',25,
 ARRAY['https://images.unsplash.com/photo-1566665797739-1674de7a421a'],
 '{"bed":"queen","view":"city"}'),

('itm_003','Presidential Suite','Premium suite with living room and luxury amenities',50000,'USD','active',2,
 ARRAY['https://images.unsplash.com/photo-1578683010236-d716f9a3f461'],
 '{"bed":"king","view":"panoramic"}'),

('itm_004','Family Room','Spacious room for families with 2 beds',18000,'USD','active',12,
 ARRAY['https://images.unsplash.com/photo-1590490360182-c33d57733427'],
 '{"beds":2,"type":"family"}'),

('itm_005','Single Room','Cozy room for solo travelers',6000,'USD','active',30,
 ARRAY['https://images.unsplash.com/photo-1618773928121-c32242e63f39'],
 '{"bed":"single"}'),

('itm_006','Twin Room','Room with two separate beds',9000,'USD','active',18,
 ARRAY['https://images.unsplash.com/photo-1582719478250-c89cae4dc85b'],
 '{"beds":2}'),

('itm_007','Executive Suite','Business suite with workspace',22000,'USD','active',6,
 ARRAY['https://images.unsplash.com/photo-1560448204-e02f11c3d0e2'],
 '{"desk":true,"wifi":"high-speed"}'),

('itm_008','Honeymoon Suite','Romantic suite for couples',30000,'USD','active',4,
 ARRAY['https://images.unsplash.com/photo-1591088398332-8a7791972843'],
 '{"theme":"romantic"}'),

('itm_009','Garden View Room','Room overlooking hotel gardens',11000,'USD','active',14,
 ARRAY['https://images.unsplash.com/photo-1551882547-ff40c63fe5fa'],
 '{"view":"garden"}'),

('itm_010','City View Room','Room with city skyline view',12000,'USD','active',20,
 ARRAY['https://images.unsplash.com/photo-1564501049412-61c2a3083791'],
 '{"view":"city"}');