package soal1

import (
	"bufio"
	"os"

	"github.com/labstack/gommon/log"
)

func Input() {
	file, err := os.Open(`input.in`)
	if err != nil {
		log.Info(err)
	}
	fOut, err := os.Create("output" + ".out")
	if err != nil {
		log.Info(err)
	}
	defer file.Close()
	defer fOut.Close()

	var scanner = bufio.NewScanner(file)
	var writer = bufio.NewWriter(fOut)
	defer writer.Flush()
	var count int
	var keyArr []string
	var valueArr []string
	for scanner.Scan() {
		var line = scanner.Text()
		if count != 0 && line != "}" {
			var key, value = KeyValue(line)
			keyArr = append(keyArr, key)
			valueArr = append(valueArr, value)
		}
		count++
	}
	log.Info(keyArr)
	log.Info(valueArr)

	for i := 0; i < len(keyArr); i++ {
		var indexPoint int
		for j := 0; j < len(keyArr[i]); j++ {
			if keyArr[i][j] == '.' {
				indexPoint = j
	}
}

func KeyValue(s string) (key, value string) {
	s = s[4:]
	if string(s[len(s)-1]) == "," {
		s = s[:len(s)-1]
	}

	var indexSlash int

	for i := 0; i < len(s); i++ {
		if string(s[i]) == ":" {
			indexSlash = i
			break
		}
		if (s[i] >= 48 && s[i] <= 57) || (s[i] >= 64 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) || string(s[i]) == "." {
			key += string(s[i])
		}
	}

	for i := indexSlash + 1; i < len(s); i++ {
		if (s[i] >= 48 && s[i] <= 57) || (s[i] >= 64 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) || string(s[i]) == "." {
			value += string(s[i])
		}
	}

	return key, value
}
