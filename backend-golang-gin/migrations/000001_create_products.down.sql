-- Hapus index terlebih dahulu
DROP INDEX IF EXISTS idx_products_category;
DROP INDEX IF EXISTS idx_products_name;

-- Hapus tabel products
DROP TABLE IF EXISTS products;