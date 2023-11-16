DROP TRIGGER set_timestamp_delete_powerbank ON tbl_powerbanks CASCADE;
DROP TRIGGER set_timestamp_delete_station ON tbl_stations CASCADE;
DROP TRIGGER set_timestamp_delete_user ON tbl_users CASCADE;
DROP FUNCTION trigger_set_timestamp_delete();

DROP TRIGGER set_timestamp_update_users ON tbl_users CASCADE;
DROP TRIGGER set_timestamp_update_powerbanks ON tbl_powerbanks CASCADE;
DROP TRIGGER set_timestamp_update_stations ON tbl_stations CASCADE;
DROP FUNCTION trigger_set_timestamp_update();

drop table tbl_station_powerbank CASCADE;
drop table tbl_user_powerbank CASCADE;
drop table tbl_addresses CASCADE;
drop table tbl_users CASCADE;
drop table tbl_powerbanks CASCADE;
drop table tbl_stations CASCADE;
drop table tbl_role CASCADE;
drop table schema_migrations CASCADE;
