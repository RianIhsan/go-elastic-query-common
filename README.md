# go-elastic-query-common

**go-elastic-query-common** adalah library Golang yang menyediakan fungsi dan utilitas umum untuk mempermudah penyusunan dan eksekusi query ke Elasticsearch.  
Library ini dirancang untuk membantu developer menyusun query Elasticsearch dengan lebih efisien dan fleksibel.

## Fitur
- **Pembuatan Query yang Mudah**: Fungsi-fungsi untuk membangun query Elasticsearch dengan cepat.
- **Reusable Components**: Komponen yang dapat digunakan kembali untuk berbagai kebutuhan query.
- **Abstraksi API**: Mempermudah interaksi dengan Elasticsearch tanpa harus menulis query secara manual.
- **Dukungan Filter dan Aggregasi**: Penyusunan filter dan aggregasi yang lebih terstruktur.

## Instalasi

Gunakan `go get` untuk menambahkan library ini ke proyek Anda:

```bash
go get github.com/RianIhsan/go-elastic-query-common
```


## How to run ?
1. Gunakan `docker-compose up -d`  untuk buat container elastic
2. Jika sudah berhasil, cek container apakah sudah active `docker ps -a`
3. Selanjutnya buka elasticvue, klik "add cluster", pilih "no auth", lalu masukan cluster name "docker-cluster" dan URL `http://localhost:9200` sesuaikan port nya dengan yang kamu pasang di `docker-compose.yml` dan sesuaikan juga di file `main.go`
4. Jika sudah bisa masuk ke cluster, silahkan jalankan server nya `go run cmd/main.go`
5. Jika terjadi error kemungkinan space memory/storage/disk kamu perlu dibersihkan
6. Jika sudah running server nya, lanjut tinggal buka postman dan test endpoint yang sudah disediakan `localhost:8080/users` dan masukan payload seperti ini

```json
   {
     "id": "1",
     "name": "John Doe",
     "email": "john.doe@example.com"
   }
```

Selamat mencoba
