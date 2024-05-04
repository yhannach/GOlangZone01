package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "unicode"
)

func main() {
    inputFile, outputFile := parseArgs()

    inputText, err := readFile(inputFile)
    if err != nil {
        fmt.Println("Error reading input file:", err)
        return
    }

    tokens := tokenize(inputText)

    modifiedTokens := applyModifications(tokens)

    err = writeFile(outputFile, modifiedTokens)
    if err != nil {
        fmt.Println("Error writing output file:", err)
        return
    }

    fmt.Println("Text modification complete.")
}

func parseArgs() (string, string) {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run main.go input_file output_file")
        os.Exit(1)
    }
    inputFile := os.Args[1]
    outputFile := os.Args[2]
    return inputFile, outputFile
}

func readFile(filename string) (string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var text string
    for scanner.Scan() {
        text += scanner.Text() + "\n"
    }
    if err := scanner.Err(); err != nil {
        return "", err
    }
    return text, nil
}

func tokenize(text string) []string {
    tokens := strings.FieldsFunc(text, func(r rune) bool {
        return !unicode.IsLetter(r) && !unicode.IsNumber(r)
    })
    return tokens
}

func applyModifications(tokens []string) []string {
    modifiedTokens := make([]string, 0)
    for i := 0; i < len(tokens); i++ {
        token := tokens[i]
        switch token {
        case "(hex)":
            modifiedTokens = replaceHexadecimal(tokens, i, modifiedTokens)
        case "(bin)":
            modifiedTokens = replaceBinary(tokens, i, modifiedTokens)
        case "(up)":
            modifiedTokens = convertUppercase(tokens, i, modifiedTokens)
        case "(low)":
            modifiedTokens = convertLowercase(tokens, i, modifiedTokens)
        case "(cap)":
            modifiedTokens = convertCapitalized(tokens, i, modifiedTokens)
        case "'":
            modifiedTokens = handleQuotation(tokens, i, modifiedTokens)
        case ".", ",", "!", "?", ":", ";":
            modifiedTokens = formatPunctuation(tokens, i, modifiedTokens)
        default:
            modifiedTokens = append(modifiedTokens, token)
        }
    }
    return modifiedTokens
}

func replaceHexadecimal(tokens []string, index int, modifiedTokens []string) []string {
    if index > 0 && isHexadecimal(tokens[index-1]) {
        decimalValue := strconv.FormatInt(parseHexadecimal(tokens[index-1]), 10)
        modifiedTokens[len(modifiedTokens)-1] = decimalValue
    }
    return modifiedTokens
}

func isHexadecimal(s string) bool {
    _, err := strconv.ParseInt(s, 16, 64)
    return err == nil
}

func parseHexadecimal(s string) int64 {
    value, _ := strconv.ParseInt(s, 16, 64)
    return value
}

func replaceBinary(tokens []string, index int, modifiedTokens []string) []string {
    return modifiedTokens
}

func convertUppercase(tokens []string, index int, modifiedTokens []string) []string {
    return modifiedTokens
}

func convertLowercase(tokens []string, index int, modifiedTokens []string) []string {
    return modifiedTokens
}

func convertCapitalized(tokens []string, index int, modifiedTokens []string) []string {
    return modifiedTokens
}

func handleQuotation(tokens []string, index int, modifiedTokens []string) []string {
    return modifiedTokens
}

func formatPunctuation(tokens []string, index int, modifiedTokens []string) []string {
    return modifiedTokens
}

func writeFile(filename string, tokens []string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    defer writer.Flush()

    for _, token := range tokens {
        _, err := writer.WriteString(token + " ")
        if err != nil {
            return err
        }
    }
    return nil
}

