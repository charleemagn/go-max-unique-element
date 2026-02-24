package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Количество элементов в m:")
	scanner.Scan()
	m, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Ошибка чтения m:", err)
		return
	}
	
	fmt.Println("Вводите элементы через пробел в m:")
	scanner.Scan()
	kStr := strings.Split(scanner.Text(), " ")
	K := make([]int, m)
	for i := 0; i < m; i++ {
		K[i], err = strconv.Atoi(kStr[i])
		if err != nil {
			fmt.Println("Ошибка чтения K:", err)
			return
		}
	}

	fmt.Println("Количество элементов в n:")
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Ошибка чтения n:", err)
		return
	}
	
	fmt.Println("Вводите элементы через пробел в n:")
	scanner.Scan()
	lStr := strings.Split(scanner.Text(), " ")
	L := make([]int, n)
	for i := 0; i < n; i++ {
		L[i], err = strconv.Atoi(lStr[i])
		if err != nil {
			fmt.Println("Ошибка чтения L:", err)
			return
		}
	}

	lSet := make(map[int]bool)
	for _, val := range L {
		lSet[val] = true
	}

	var uniqueK []int
	for _, val := range K {
		if !lSet[val] {
			uniqueK = append(uniqueK, val)
		}
	}

	if len(uniqueK) == 0 {
		fmt.Println("Нет элементов в K, отсутствующих в L")
		return
	}

	sort.Sort(sort.Reverse(sort.IntSlice(uniqueK)))
	fmt.Printf("Наибольший элемент K, отсутствующий в L: %d\n", uniqueK[0])
}
