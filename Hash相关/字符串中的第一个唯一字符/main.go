package main

func firstUniqChar(s string) int {
	hashTable := make(map[byte]int)
	for _, v := range s {
		hashTable[byte(v)]++
	}
	for k, v := range s {
		if hashTable[byte(v)] == 1 {
			return k
		}
	}
	return -1
}

func main() {

}
