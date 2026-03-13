    Client (Postman/Browser)
         │
         ▼
    main.go          ← titik awal, jalankan server
         │
         ▼
  endpoint.go        ← daftar semua URL/route
  (rest/)            contoh: GET /users, POST /users
         │
         ▼
  handler.go         ← terima request, validasi, kirim response
  (handler/)         
         │
         ▼
  crud/*.go          ← logika data (add, get, update, delete)
  (crud/)            
         │
         ▼
    db/              ← koneksi & data tersimpan di sini




    package rest

import (
    "net/http"
    "<module_name>/handler"  // ← hubungkan ke handler
)

// Clue: buat fungsi yang mendaftarkan semua endpoint
func SetupRoutes() {
    http.HandleFunc("/users",    handler.HandleUser)
    http.HandleFunc("/users/add", handler.HandleAdd)
    // tambah route lain...
}


package handler

import (
    "encoding/json"  // ← untuk baca/tulis JSON
    "net/http"
    "<module_name>/crud"  // ← hubungkan ke crud
)

// Clue: cek method dulu, baru arahkan ke fungsi yang tepat
func HandleUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        // panggil crud.GetData()
    } else if r.Method == "POST" {
        // panggil crud.AddData()
    }
}

package db

// Clue: karena ini latihan, simpan data di memori dulu (slice)
// Bukan database sungguhan dulu!

type User struct {
    ID    int    `json:"id"`
    Nama  string `json:"nama"`
    Email string `json:"email"`
}

// ini "database" sementara di memori
var DataUsers = []User{}


package crud

import "<module_name>/db"

// Clue: tambahkan data baru ke slice
func AddData(user db.User) {
    db.DataUsers = append(db.DataUsers, user)
}
```

---

## 🗺️ Urutan Belajar yang Disarankan
```
Langkah 1 → Isi db/        (buat struct & variabel data)
     ↓
Langkah 2 → Isi crud/      (fungsi get, add, delete, update)
     ↓
Langkah 3 → Isi handler/   (terima request, panggil crud)
     ↓
Langkah 4 → Isi endpoint/  (daftarkan route)
     ↓
Langkah 5 → Isi main.go    (jalankan server)
     ↓
Langkah 6 → Test pakai Postman! 🎉



package crud

import "<module_name>/db"

// Clue: kembalikan slice data dari db
func GetData() []db.User {
    return db.DataUsers
}