package yr

import (
	"testing"
	"os"
	"bufio"
)

func TestCelsiusToFahrenheitString(t *testing.T) {
     		type test struct {
		input string
		want string
     	}
     		tests := []test{
	     	{input: "6", want: "42.8"},
	     	{input: "0", want: "32.0"},
		{input: "-11", want: "12.2"},
	      }

     	for _, tc := range tests {
	     	got, _ := CelsiusToFahrenheitString(tc.input)
	     	if !(tc.want == got) {
		     	t.Errorf("expected %s, got: %s", tc.want, got)
	     		}
     		}
	}

	// Forutsetter at vi kjenner strukturen i filen og denne implementasjon 
	// er kun for filer som inneholder linjer hvor det fjerde element
	// på linjen er verdien for temperatrmaaling i grader celsius
func TestCelsiusToFahrenheitLine(t *testing.T) {
     		type test struct {
		input string
		want string
     	}
     		tests := []test{
	     	{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
	     	{input: "Kjevik;SN39040;18.03.2022 01:50;0", want: "Kjevik;SN39040;18.03.2022 01:50;32.0"},
		{input: "Kjevik;SN39040;18.03.2022 01:50;-11", want: "Kjevik;SN39040;18.03.2022 01:50;12.2"},
		//{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;", want: "Data er basert på gyldig data (per 18.03.2023) (CCBY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Tj12501"},

     	}

     	for _, tc := range tests {
	     	got, _ := CelsiusToFahrenheitLine(tc.input)
	     	if !(tc.want == got) {
		     	t.Errorf("expected %s, got: %s", tc.want, got)
	     		}
     		}	
	}

func TestLineCount(t *testing.T) {
    	file, err := os.Open("../kjevik-temp-fahr-20220318-20230318.csv")
    	if err != nil {
        	t.Errorf("Failed to open file: %v", err)
    	}
    	defer file.Close()

    	scanner := bufio.NewScanner(file)
    	lineCount := 0
    	for scanner.Scan() {
        	lineCount++
    	}

    	if lineCount != 16756 {
        	t.Errorf("Unexpected line count: %d", lineCount)
    		}
	}

func TestLastLine(t *testing.T) {
    file, err := os.Open("../kjevik-temp-fahr-20220318-20230318.csv")
    if err != nil {
        t.Fatalf("failed to open file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lastLine string
    for scanner.Scan() {
        lastLine = scanner.Text()
    }

    expectedLastLine := "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Tj12501"
    if lastLine != expectedLastLine {
        t.Errorf("Last line is incorrect: got %q, want %q", lastLine, expectedLastLine)
    }
}

func TestAverageTemperature(t *testing.T) {
     // Test a file with valid data
	avg, err := CalculateAverageTemperature("../kjevik-temp-celsius-20220318-20230318.csv", "c")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if avg <= 8.55 || avg >= 8.57 {
		t.Errorf("expected average temperature 8.56, got %f", avg)
	}
}
