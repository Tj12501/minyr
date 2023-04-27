package yr

import (
        "fmt"
        "bufio"
        "os"
        "io"
        "strconv"
        "strings"
        "errors"
        "github.com/Tj12501/funtemps/conv"
        //"minyr/conv"
        //"log"
)
	func CelsiusToFahrenheitString(celsius string) (string, error) {
        	var fahrFloat float64
        	var err error
        	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
                	fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
        	}
        	fahrString := fmt.Sprintf("%.1f", fahrFloat)
        	return fahrString, err
	}

	// Forutsetter at vi kjenner strukturen i filen og denne implementasjon
	// er kun for filer som inneholder linjer hvor det fjerde element
	// på linjen er verdien for temperaturaaling i grader celsius
	func CelsiusToFahrenheitLine(line string) (string, error) {

        	dividedString := strings.Split(line, ";")
        	var err error

	        if (len(dividedString) == 4) {
                	dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
                	if err != nil {
                        	return "", err
                	}
        	} else {
                	return "", errors.New("linje har ikke forventet format")
        	}
        	return strings.Join(dividedString, ";"), nil

        	//return "Kjevik;SN39040;18.03.2022 01:50;42.8", err
	}



	func GetNumberOfLines(filename string) (int, error) {
	    file, err := os.Open(filename)
	    if err != nil {
	        return 0, err
	    }
	    defer file.Close()

	    reader := bufio.NewReader(file)

	    count := 0
	    for {
	        _, err := reader.ReadString('\n')
	        if err == io.EOF {
	            break
	        } else if err != nil {
 	           return 0, err
	        }

	        count++
	    }

	    return count, nil
	}

	func GetLastLine(filename string) (string, error) {
        	file, err := os.Open(filename)
	        if err != nil {
	                return "", err
	        }
        	defer file.Close()

        	var lastLine string
        	scanner := bufio.NewScanner(file)
        	for scanner.Scan() {
                	lastLine = scanner.Text()
        	}

        	if err := scanner.Err(); err != nil {
                	return "", err
        	}

        	return lastLine, nil
}
