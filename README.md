# go-datatable
Go Datatable örnek uygulaması ile jQuery Datatable paketinin Go ile sunucu taraflı (serverside) çalışmasını göstermek için hazırlanmıştır.

## Gereksinimler
- Go
- Gorm
- Gin
- Postgres

## Kurulum
- Öncelikle github paketini bilgisayarınızdaki Go dizinine (örn. $HOME/go/src/) indiriniz.
- Paket içindeki backend/inc/database.go dosyasindaki DATABASE INFORMATION bloğuna Postgres veritabanı bilgilerinizi giriniz.
- Terminalden paket içindeki backend/ dizinine girerek "go run main.go" komutu ile uygulamayı başlatınız.
- Backend yani Go uygulaması çalışmaya başladıktan sonra proje ana dizinindeki index.html dosyasını herhangi bir tarayıcıyla açınız.


## Kullanılan Paketler ve Teknoloji Açıklamaları
- jQuery Datatable (https://www.datatables.net/)
Projenin temelini oluşturan jQquery eklentisi

- Gin (https://github.com/gin-gonic/gin)
Popüler Go API framework'ü.

- Gorm (https://gorm.io/)
ORM veritabanı yönetim paketi.

- Postgres (https://gorm.io/docs/connecting_to_the_database.html)
Veritabanı olarak Postgres kullanılmıştır. Gorm paketi ile gerekli bağlantılar yapılmıştır. Gorm'un esnek yapısı ile proje rahatlıkla SqlLite veya MySQL'e dönüştürülebilir.

## Server Tarafında Yapılan İşlemler
- Bütün sütunlar için Search
- Pagination
- Query limit
- Query order

![Go Datatable Screenshot](https://github.com/yakuter/go-datatable/blob/master/screenshot.png)