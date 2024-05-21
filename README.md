# Minder
Min Tinder, Minder For App


## Struktur

Ada beberapa struktur utama,

- `helper`: helper adalah kode yang berisi fungsi untuk digunakan berulang kali dan efisien

- `config`: config berisi fungsi untuk melakukan sesuatu untuk pertama kali, misalnya menghubungkan ke database dan menjalankannya secara singleton
- `/src/delivery`: menghubungkan logika dengan presenter, presenter yang digunakan bisa rest json, grpc, dalam hal ini menggunakan rest json
- `src/model` : representasi data
- `src/repository` ; fungsi untuk melakukan handle dengan database
- `src/usecase` : dimana akan menghandle terkait logic bussiness



## Cara Menjalankannya

Requirement
- `docker` 
- `make`


## Instalasi
- Clone repo ini
- jalankan `make up`, env db bisa check di .env
- jalankan `go run main.go`



## What If

- Database belum termigrate ? jalankan make run-migrate
- Bentrok Port? jalankan  `docker system prune` 
- menonaktifkan server database ? jalankan  `docker-componse down` 

