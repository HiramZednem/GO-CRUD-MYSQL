# How to Run this Project!
Hello Stranger, or HiramZednem from the future, if you want to compile this project you need to have created a database of mysql running in your computer or in aws, google cloud or a cloud service.
## Step 1:
Create a file named `.env` in the root of the project
## Step 2:
Add the DSN variable like this:
### Postgresql
> DSN="host=localhost user=gorm password=gorm dbname=gorm port=5432
### Mysql
> DSN="user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

## Step 3
Now that the file .env was created, you are able to execute the command
> go run .

Or if you have `cosmtrek/air`, use
> air 