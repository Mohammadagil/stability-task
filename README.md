# Stability Team Technical Test

This repository contains a simple Task Manager API built with Go and Fiber.

## Setup

Install dependencies:
go mod tidy

Run the server:
go run main.go

Server will run at:
http://localhost:3000

## Available Endpoints

GET /tasks  
GET /tasks/:id  
POST /tasks  
PUT /tasks/:id    
DELETE /tasks/:id

## Issues yang Saya Temukan

1.	Tidak ada validasi id dan title. Lokasi, task_handler.go.
2.	Status code salah untuk “task not found”. Lokasi, task_handler.go->GetTask().
3.	Error handling id diabaikan. Lokasi, task_handler.go->GetTask() dan task_handler.go->DeleteTask()
4.	Id task tidak auto-generate. Lokasi, task_handler.go->CreateTasks()
5.	Delete task selalu sukses meskipun id tidak ada. Lokasi, task_handler.go->DeleteTask()
6.	Pointer bug di GetTaskByID, mengubah satu task akan mengubah task lain. Lokasi, task_store.go->GetTaskByID().
7.	Missing return di DeleteTask. Lokasi, task_store.go->DeleteTask().

## Perbaikan Issues yang Saya Lakukan
1.	Menambahkan pemeriksaan di setiap titik yang membutuhkan validasi. Cek apakah id kosong, apakah id benar-benar angka, apakah id negative. Cek apakah title kosong, apakah title hanya berisi spasi, apakah Panjang title maksimal 100 karakter.
2.	Saya mengubah kode status dari 200 menjadi 404 ketika tugas tidak ditemukan (sesuai standar).
3.	Saya menambahkan pengecekan untuk memastikan ID benar-benar angka. Jika ternyata bukan angka, sistem akan langsung merespon dengan response : StatusBadRequest.
4.	Saya membuat sistem yang secara otomatis menghasilkan ID baru setiap kali pengguna membuat tugas. Caranya:
•	Cari ID terbesar yang sudah ada (misalnya ID 5 adalah yang terbesar)
•	Tambahkan 1 (menjadi 6)
•	Gunakan ID baru itu untuk tugas yang akan dibuat
5.	Saya menambahkan pengecekan sebelum menghapus:
•	Cek dulu apakah tugas dengan ID tersebut benar-benar ada
•	Jika tidak ada → kirim response : StatusNotFound dengan pesan "task not found"
•	Jika ada → baru hapus dan kirim pesan sukses
6.	Saya mengubah cara sistem mengambil data tugas, saya mengambil langsung data asli dari kumpulan data. Jadi setiap tugas memiliki alamat memorinya sendiri-sendiri. Saat mengubah tugas id 1 maka tugas dengan id 2 aman tidak terpengaruh. 
7.	Saya menambahkan perintah "return” setelah proses hapus selesai. Begitu tugas ditemukan dan dihapus, sistem langsung berhenti mencari.

## Penambahan yang Saya Lakukan 
1.	Menambahkan fungsi/endpoint baru untuk Update task. 
Lokasi, task_handler.go->UpdateTask(). Melakukan validasi id, periksa apakah task ada, melakukan parse request body, melakukan update jika ada title yang dikirim, melakukan validasi title, setelah semua terpenuhi data baru akan tersimpan.
3. Menambah logging midleware untuk mencatat semua aktivitas request yang masuk ke server, etiap kali ada yang memanggil API

## Melakukan Testing API Menggunakan Postman 
Buat Collection Baru

Buat 5 request testing API :
1.	All Tasks, method : GET, url : http://localhost:3000/tasks
2.	Task by ID, method : GET, url : http://localhost:3000/tasks/{id}
3.	Create Task, method : POST, url : http://localhost:3000/tasks    
    Body->raw->JSON    
{      
“title”:    
“done”: (opsional)    
}      
5.	Update Task, method : PUT, url : http://localhost:3000/tasks/{id}    
    Body->raw->JSON    
{    
“title”:    
“done”: (opsional)    
}    
7.	Delete Task, method : DELETE, url : http://localhost:3000/tasks/{id}
