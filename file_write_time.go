package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	// Получаем текущее время
	currentTime := time.Now()

	// Форматируем время и дату в строки
	timeString := currentTime.Format("15:04:05") // Часы:Минуты:Секунды
	dateString := currentTime.Format("2006-01-02") // Год-месяц-день
	dayOfWeek := currentTime.Weekday()// День недели

	// Определяем путь к рабочему столу
	desktopPath := getDesktopPath()
	filePath := filepath.Join(desktopPath, "file_write_time.txt") // Имя файла

	// Создаем или открываем файл для записи
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close() // Закрываем файл в конце

	// Записываем текущее время, день недели и дату в файл в столбик для читаемости
	_, err = file.WriteString(fmt.Sprintf("Текущее время: %s\nДень недели: %s\nДата: %s\n", timeString, dayOfWeek, dateString))
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	// Выводим сообщение в терминал
	//fmt.Println("Информация записана в файл:", filePath)
	fmt.Println("SEE THE DESKTOP :)")
}

// getDesktopPath возвращает путь к рабочему столу пользователя
func getDesktopPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Ошибка при получении домашнего каталога:", err)
		return ""
	}

	// Определяем путь к рабочему столу в зависимости от операционной системы
	if runtime.GOOS == "windows" {
		return filepath.Join(homeDir, "Desktop")
	} else if runtime.GOOS == "darwin" { // macOS
		return filepath.Join(homeDir, "Desktop")
	} else { // Linux и другие
		return filepath.Join(homeDir, "Рабочий стол")
	}
}

