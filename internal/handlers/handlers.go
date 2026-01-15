package handlers

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// возвращает HTML из файла index.html
func HandleMain(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	// парсинг html-формы из файла index.html
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "ошибка при парсинге формы", http.StatusInternalServerError)
		return
	}

	// получение файла из формы
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// чтение данных из файла
	root, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	// передача этих данных в функцию автоопределения из пакета service
	result := service.Converter(string(root))

	// создание файла
	filename := time.Now().UTC().Format("20060102150405") + filepath.Ext(header.Filename)
	outFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// запись в файл
	_, err = outFile.Write([]byte(result))
	if err != nil {
		http.Error(w, "ошибка при записи файла", http.StatusInternalServerError)
		return
	}

	// возврат результата пользователю
	w.Write([]byte(result))
}
