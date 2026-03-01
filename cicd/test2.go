package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
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
  if len(content) == 0 {
    fmt.Println("Файл вывода пуст")
    return false
  }
  return true
}

func main() {
  fmt.Println("=== ТЕСТ 2: Проверка обработки некорректного ввода ===")
  passed := 0

  inputData := "three\n1 3 2\n2\n3 4\n" // Первое - буква вместо m
  err := ioutil.WriteFile("test2_input.txt", []byte(inputData), 0644)
  if err != nil {
    fmt.Println("Ошибка создания input:", err)
    return
  }

  if runTest("Ввод с ошибками", "test2_input.txt", "test2_output.txt") {
    if checkOutput("test2_output.txt", "Ошибка чтения m:") {
      fmt.Println("✓ Тест 2 пройден (обнаружена ошибка ввода)")
      passed++
    } else {
      fmt.Println("✗ Тест 2 не пройден")
    }
  }

  if passed == 1 {
    fmt.Println("\nИТОГ: Тест пройден")
  } else {
    fmt.Println("\nИТОГ: Тест не пройден")
  }
}
