package main

import (
	"log"
	"os"

	"github.com/Julianarwansah/be-catalog-p4-1123150112/config"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan environment system")
	}

	// 2. Inisialisasi Database (AutoMigrate terpanggil di sini)
	config.InitDatabase()

	// 3. Inisialisasi Firebase (Opsional jika file JSON belum ada)
	credPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if _, err := os.Stat(credPath); err == nil {
		config.InitFirebase()
	} else {
		log.Printf("Info: Firebase credentials tidak ditemukan di %s, skip...", credPath)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server siap! Silakan cek phpMyAdmin di MAMP untuk tabel baru.")
	log.Printf("Aplikasi berjalan di port %s", port)
}
