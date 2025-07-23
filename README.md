# Aplikasi Opsi Pembayaran

Layanan web sederhana yang dibuat dengan Go untuk menyimulasikan pengambilan beberapa metode pembayaran secara bersamaan menggunakan goroutine.

## Alur Kerja

Sistem Opsi Pembayaran ini akan menampilkan daftar metode pembayaran yang tersedia, yang meliputi:

- OVO
- DANA
- GoPay
- ShopeePay
- OneKlik
- BRIDD
- LinkAja

Ketika _endpoint_ untuk daftar metode pembayaran diakses, _backend_ akan memanggil API profil untuk setiap mitra pembayaran. Permintaan ini dieksekusi secara paralel untuk efisiensi yang lebih baik.

Setelah semua respons dari mitra diterima, _backend_ akan mengumpulkan data yang diperlukan dan mengembalikan daftar lengkap metode pembayaran beserta informasi profil yang relevan ke _frontend_.

## Rincian Perubahan

Berikut adalah rincian perubahan dan penambahan yang telah saya lakukan pada proyek ini:

1.  **Saya menambahkan "LinkAja" sebagai opsi pembayaran baru.** Penambahan ini mengikuti pola yang sama dengan metode pembayaran lainnya, di mana pemanggilan API disimulasikan dan dijalankan secara bersamaan dalam goroutine terpisah untuk menjaga efisiensi aplikasi.
2.  **Saya mengoptimalkan logika konkurensi.** Saya memperbaiki cara kerja goroutine agar pemanggilan API yang lambat dieksekusi terlebih dahulu sebelum mengunci data. Perubahan ini memastikan bahwa semua panggilan benar-benar berjalan secara paralel, yang secara signifikan meningkatkan performa dan kecepatan respons aplikasi.

## Setup Proyek

### Koneksi Database

Proyek ini dikonfigurasi untuk terhubung ke dua database MySQL saat startup:

- **Database 1:** `produk`
- **Database 2:** `payment`

Konfigurasi koneksi dapat ditemukan di `internal/repository/db.go`. Saat ini, koneksi database ini diinisialisasi tetapi belum digunakan dalam alur logika untuk mengambil opsi pembayaran, yang masih menggunakan data simulasi.

### Menjalankan Aplikasi

```bash
go mod init payment-options
go mod tidy
go run cmd/main.go
```

Server akan berjalan di `http://localhost:8081`

## Menguji Endpoint

Anda dapat menguji menggunakan `curl`:

```bash
curl http://localhost:8081/payment/options
```

### Contoh Respons

```json
{
  "returnCode": "200",
  "returnDesc": "success",
  "data": {
    "ovo": {
      "account": "6288xx",
      "status": "Active",
      "balance": "10000",
      "icon": "https://sampleurl.com/ovo.jpg"
    },
    "dana": {
      "account": "6288xx",
      "status": "Active",
      "balance": "10000",
      "icon": "https://sampleurl.com/dana.jpg"
    },
    ...
  }
}
```

## Lisensi

MIT
