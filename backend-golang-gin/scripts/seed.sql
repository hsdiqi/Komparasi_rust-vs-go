TRUNCATE TABLE products RESTART IDENTITY;
INSERT INTO products (name, description, price, stock, category, created_at, updated_at)
SELECT
  'Product ' || gs,
  'Description for product ' || gs,
  ((gs % 1000) + 1) * 1000,
  (gs % 500),
  CASE gs % 10
    WHEN 0 THEN 'phone'
    WHEN 1 THEN 'laptop'
    WHEN 2 THEN 'tablet'
    WHEN 3 THEN 'accessory'
    WHEN 4 THEN 'camera'
    WHEN 5 THEN 'audio'
    WHEN 6 THEN 'monitor'
    WHEN 7 THEN 'keyboard'
    WHEN 8 THEN 'mouse'
    ELSE 'printer'
  END,
  NOW(), NOW()
FROM generate_series(1, :COUNT) AS gs;
