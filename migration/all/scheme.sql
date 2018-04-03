CREATE DATABASE circle owner postgres encoding 'utf8';



/* Drop Tables */
DROP TABLE IF EXISTS user_app;


CREATE TABLE user_app
(
	id_user_app serial NOT NULL UNIQUE,
	uuid char(36) NOT NULL UNIQUE,
	-- 文字数制限はユーザ登録要件未確定のため仮の値です。
	user_name varchar(20),
	user_profile_image_url varchar(256),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_user_app)
) WITHOUT OIDS;

ALTER SEQUENCE user_app_id_user_app_SEQ INCREMENT 1 RESTART 1;