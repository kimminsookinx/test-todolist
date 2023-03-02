#!/bin/bash
#
#mysql variables
#
todo_db_name='todo5'
todo_user_name='create05'
todo_user_pw='create05'
todo_table_name='item'
sql_file_name='todo.sql'
#
#create database
#
echo "CREATE DATABASE IF NOT EXISTS \`$todo_db_name\` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT ENCRYPTION='N';" >> $sql_file_name
#
#create user and grant privlege to db
#
echo "CREATE USER IF NOT EXISTS '$todo_user_name'@'%' IDENTIFIED BY '$todo_user_pw';" >> $sql_file_name
echo "GRANT ALL ON \`$todo_db_name\`.* to '$todo_user_name'@'%';" >> $sql_file_name
#
#create table
#
echo "USE \`$todo_db_name\`;" >> $sql_file_name
echo "CREATE TABLE $todo_table_name (" >> $sql_file_name
echo "\`id\` int NOT NULL AUTO_INCREMENT, " >> $sql_file_name
echo "\`description\` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL, " >> $sql_file_name
echo "\`created_at\` datetime DEFAULT CURRENT_TIMESTAMP, " >> $sql_file_name
echo "\`last_updated_at\` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, " >> $sql_file_name
echo "\`done\` tinyint DEFAULT '0', " >> $sql_file_name
echo "\`deleted\` tinyint DEFAULT '0', " >> $sql_file_name
echo "\`deleted_at\` datetime DEFAULT NULL, " >> $sql_file_name
echo "PRIMARY KEY (\`id\`) " >> $sql_file_name
echo ") AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci; " >> $sql_file_name
#
#add initial data
#
echo "INSERT INTO \`$todo_db_name\`.\`$todo_table_name\`(description) VALUES(\"first todo item\"); " >> $sql_file_name
echo "SELECT SLEEP(3);" >> $sql_file_name
echo "INSERT INTO \`$todo_db_name\`.\`$todo_table_name\`(description) VALUES(\"second todo item\"); " >> $sql_file_name