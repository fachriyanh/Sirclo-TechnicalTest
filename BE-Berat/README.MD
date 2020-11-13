#How to Run The Program
This Application will let you do CRUD data from web

1. Configure your Gopath and make sure you clone this project on the right directory

2. Install sql library
    `go get "github.com/go-sql-driver/mysql"`

3. Prepare your database

4. Add tableberat
    ```
    CREATE TABLE tableberat (
    tanggal VARCHAR(20) PRIMARY KEY,
    max INT(30),
    min INT(30),
    perbedaan INT(30)
    )
    ```
    you can also custom the configuration at config.go

5. Run your web server (Apache)

6. Run the backend
    'go run main.go`

7. Open the site through the browser
    `localhost:3306`

#Documentation

Show Data
![showdata](https://user-images.githubusercontent.com/43384324/99107722-dcc68380-2618-11eb-86cc-58f9965afb76.PNG)

In this page you can see the maximum, minimum, and different between them for each date. You can also find the average value for maximum, minimum and the different between them.

Add and Edit Data
![add data](https://user-images.githubusercontent.com/43384324/99107187-0af79380-2618-11eb-983a-f2c429c3ef1e.PNG)

In this page you can add new data by choosing the date and write maximum and minimum number. If you choose the exist date in database, so you will edit the data on that date.

Detail Data
![detaildata](https://user-images.githubusercontent.com/43384324/99107231-13e86500-2618-11eb-9056-b1ddf163a2cc.PNG)

When you click the date in the table you will go directly to detail information of that date.