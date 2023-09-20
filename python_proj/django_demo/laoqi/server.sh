#!/bin/bash

case $1 in
  "run") python manage.py runserver
  ;;
  "sqlite") sqlite3 db.sqlite3
  ;;
esac