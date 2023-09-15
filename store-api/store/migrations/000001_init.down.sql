DROP TABLE IF EXISTS purchase;
DROP PROCEDURE IF EXISTS register_purchase;

DROP TABLE IF EXISTS products;
DROP PROCEDURE IF EXISTS create_product;
drop function if exists get_products_page
;
drop function if exists search_product
;
drop function if exists get_product_by_path
;

DROP TABLE IF EXISTS sessions;
DROP PROCEDURE IF EXISTS create_session;
drop function if exists update_session
;
DROP PROCEDURE IF EXISTS delete_expired_sessions;

DROP TABLE IF EXISTS users;
DROP PROCEDURE IF EXISTS insert_user;
drop function if exists get_user_by_email_or_username
;
drop function if exists get_user_by_id
;
drop function if exists get_user_by_username
;
