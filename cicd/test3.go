package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
  "strings"
)

func runTest(testName, inputFile, outputFile string) bool {
  fmt.Printf("\n=== Запуск теста: %s ===\n", testName)
  cmd := exec.Command("../usr/bin/max_unique_element")
  inFile, err := os.Open(inputFile)
  if err != nil {
    fmt.Println("Ошибка открытия input:", err)
    return false
  }
  defer inFile.Close()
  cmd.Stdin = inFile

  outFile, err := os.Create(outputFile)
  if err != nil {
    fmt.Println("Ошибка создания output:", err)
    return false
  }
  defer outFile.Close()
  cmd.Stdout = outFile

  err = cmd.Run()
  if err != nil {
    fmt.Println("Ошибка выполнения программы:", err)
    return false
  }
  fmt.Println("Программа выполнилась успешно")
  return true
}

func checkOutput(outputFile, expectedText string) bool {
  data, err := ioutil.ReadFile(outputFile)
  if err != nil {
    fmt.Println("Не удалось открыть файл вывода:", err)
    return false
  }
  content := string(data)
  return strings.Contains(content, expectedText)
}

func main() {
  fmt.Println("=== ТЕСТ 3: Проверка точного соответствия вывода ===")
  passed := 0

  inputData := "5\n10 20 30 40 50\n3\n20 40 60\n" // K=[10,20,30,40,50], L=[20,40,60] → unique [10,30,50], max 50
  err := ioutil.WriteFile("test3_input.txt", []byte(inputData), 0644)
  if err != nil {
    fmt.Println("Ошибка создания input:", err)
    return
  }

  if runTest("Точный ввод", "test3_input.txt", "test3_output.txt") {
    if checkOutput("test3_output.txt", "Наибольший элемент K, отсутствующий в L: 50") {
      fmt.Println("✓ Тест 3 пройден (вывод соответствует ожидаемому)")
      passed++
    } else {
      fmt.Println("✗ Тест 3 не пройден (вывод не совпадает)")
    }
  }

  if passed == 1 {
    fmt.Println("\nИТОГ: Тест пройден")
  } else {
    fmt.Println("\nИТОГ: Тест не пройден")
  }
}
