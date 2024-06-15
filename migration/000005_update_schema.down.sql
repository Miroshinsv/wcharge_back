DROP TRIGGER set_timestamp_delete_powerbank ON powerbanks CASCADE;
DROP TRIGGER set_timestamp_delete_station ON stations CASCADE;
DROP TRIGGER set_timestamp_delete_user ON users CASCADE;
DROP FUNCTION trigger_set_timestamp_delete();

DROP TRIGGER set_timestamp_update_users ON users CASCADE;
DROP TRIGGER set_timestamp_update_powerbanks ON powerbanks CASCADE;
DROP TRIGGER set_timestamp_update_stations ON stations CASCADE;
DROP FUNCTION trigger_set_timestamp_update();

drop table rel__stations__powerbanks CASCADE;
drop table rel__users__powerbanks CASCADE;
drop table addresses CASCADE;
drop table users CASCADE;
drop table powerbanks CASCADE;
drop table stations CASCADE;
drop table role CASCADE;
drop table schema_migrations CASCADE;
