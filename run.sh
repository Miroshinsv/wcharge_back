#!/bin/sh

migrate_flag=`cat migrate_flag`

echo migrate_flag

if [ $migrate_flag -eq '1' ]
then
  echo "Start ./wcharge_back"
  echo 1 > migrate_flag
  ./wcharge_back
else
  echo "Start ./wcharge_back_migrate"
    echo 1 > migrate_flag
    ./wcharge_back_migrate
fi