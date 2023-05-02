package main
import (
        "log"
        "os"
        "bufio"
        "fmt"
        //"strconv"
        "strings"
        "github.com/Tj12501/minyr/yr"
	//"minyr/yr"
)

const mainFile = "kjevik-temp-celsius-20220318-20230318.csv"
const destinationFile = "kjevik-temp-fahr-20220318-20230318.csv"

func main() {
    choice := differentOptions()

    if choice == "exit" {
        fmt.Println("Exit")
        os.Exit(0)
    } else if choice == "convert" {
        if err := convertOption(); err != nil {
            log.Fatal(err)
        }
    } else if choice == "average" {
        if err := averageOption(); err != nil {
            log.Fatal(err)
        }
    } else {
        fmt.Println("This is not an option")
    }
}

func differentOptions() string {
    fmt.Print("Do you want to 'convert' the temperatures or calculate the 'average' temperature?\n'convert' will convert all temperatures to Fahrenheit.\n'average' will calculate the average temperature for the whole period.\nEnter 'convert' or 'average' to continue, or 'exit' to quit: ")
        choice, err := func() (string, error) {
        reader := bufio.NewReader(os.Stdin)
        choice, err := reader.ReadString('\n')
        return strings.TrimSpace(strings.ToLower(choice)), err
}()

        if err != nil {
                log.Fatal(err)
        }


        return choice
}
func convertOption() error {
    // Checks if the destination file already exists
    if _, err := os.Stat(destinationFile); err == nil {
        // Destination file exists
        scanner := bufio.NewScanner(os.Stdin)
        for {
            fmt.Print("The file already exists, do you want to generate it again? (y/n): ")
            scanner.Scan()
            confirm := strings.TrimSpace(strings.ToLower(scanner.Text()))
            if confirm == "y" || confirm == "yes" {
                // Generates the destination file again
                if err := generateOutputFile(); err != nil {
                    return err
                }
                fmt.Println("File generated")
                break
            } else if confirm == "n" || confirm == "no" {
                // Does not generate the destination file again
                fmt.Println("Exiting program")
                break
            } else {
                fmt.Println("Invalid input, please try again.")
            }
        }
    } else {
        // Destination file does not exist and the program generates a new file
        if err := generateOutputFile(); err != nil {
            return err
        }
        fmt.Println("File generated")
    }
    return nil
}

func averageOption() error {
    fmt.Print("In which unit of measurement do you want the average temperature? (c/f): ")

    reader := bufio.NewReader(os.Stdin)
    unit, err := reader.ReadString('\n')
    if err != nil {
        return err
    }
    unit = strings.TrimSpace(strings.ToLower(unit))

    if unit == "c" {
        avgTemp, err := yr.CalculateAverageTemperature(mainFile, "c")
        if err != nil {
            return err
        }
        fmt.Printf("Average temperature: %.2f °C\n", avgTemp)
    } else if unit == "f" {
        avgTemp, err := yr.CalculateAverageTemperature(destinationFile, "f")
        if err != nil {
            return err
        }
        fmt.Printf("Average temperature: %.2f °F\n", avgTemp)
    } else {
        fmt.Println("Invalid unit of measurement, choose between c and f")
    }

    return nil
}

func generateOutputFile() error {
        mainFile, err := os.Open(mainFile)
        if err != nil {
                return err
        }
        defer mainFile.Close()

        destinationFile, err := os.Create(destinationFile)
        if err != nil {
                return err
        }
        defer destinationFile.Close()

        scanner := bufio.NewScanner(mainFile)
        writer := bufio.NewWriter(destinationFile)
        defer writer.Flush()

        // get number of lines in file
        totalLines, err := yr.GetNumberOfLines(mainFile.Name())
        if err != nil {
                fmt.Println("Error while reading lines:", err)
        } else {
                fmt.Println("The total number of lines is:", totalLines)
        }

        lineCount := 0
        for scanner.Scan() {
                lineCount++
                line := scanner.Text()
                if lineCount == 1 {
                        // Write the first line to the output file as is
                        _, err = writer.WriteString(line + "\n")
                        if err != nil {
                                return err
                        }
                        continue
                }

                // Process the line (convert temperature and format output)
                processedLine, err := yr.CelsiusToFahrenheitLine(line)
                if err != nil {
                        return err
                }

                if lineCount < totalLines {
                        // Write processed line to output file
                        _, err = writer.WriteString(processedLine + "\n")
                        if err != nil {
                                return err
                        }
                } else {
                        // Write test string for the last line
                        _, err = writer.WriteString("Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Tj12501")
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}
