DROP TRIGGER IF EXISTS items_updated_at ON items;
DROP TRIGGER IF EXISTS items_search_vec ON items;
DROP FUNCTION IF EXISTS items_search_vector_update();
DROP TABLE IF EXISTS items;