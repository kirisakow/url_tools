package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Reads input either from stdin or from the argument, if any.
// Returns the input as a string.
func read_input() []string {
    var urls_to_clean []string
    stat, _ := os.Stdin.Stat()
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        // Read URLs from stdin
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            urls_to_clean = append(urls_to_clean, scanner.Text())
        }
    } else {
        // Read URLs from arguments
        if len(os.Args) > 1 {
            urls_to_clean = os.Args[1:]
        }
    }
    // Keep only non-empty strings
    var tmp []string
    for _, n := range urls_to_clean {
        if strings.TrimSpace(n) != "" {
            tmp = append(tmp, n)
        }
    }
    urls_to_clean = tmp
    return urls_to_clean
}

// Reads the unwanted_params.txt file and returns its content as an array of strings.
func get_unwanted_params(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var unwanted_params []string
	for scanner.Scan() {
		unwanted_params = append(unwanted_params, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return unwanted_params, nil
}

// Removes the unwanted query parameter from the URL.
func clean_url(url_to_clean, unwanted_param string) string {
    // If the unwanted param contains an asterisk, replace it with a regex pattern
    unwanted_param = strings.ReplaceAll(unwanted_param, "*", "[^&=?]*")
    // Use a regular expression to match the unwanted query parameter and remove it
    r := regexp.MustCompile(fmt.Sprintf(`[&]?%s=[^&]*`, unwanted_param))
    url_to_clean = r.ReplaceAllString(url_to_clean, "")
    // remove final ? symbol if no params left
    url_to_clean = regexp.MustCompile("[?]$").ReplaceAllString(url_to_clean, "")
    return url_to_clean
}

func main() {
    // Read input either from stdin or from the argument, if any
    urls_to_clean := read_input()

    // Read the unwanted_params.txt file and get the list of unwanted params
    unwanted_params, _ := get_unwanted_params("unwanted_params.txt")

    // Go through the list of URLs, then through that of unwanted params
    for _, url_to_clean := range urls_to_clean {
        for _, unwanted_param := range unwanted_params {
            // Check if the unwanted param does not contain a '?' symbol
            if !strings.Contains(unwanted_param, "?") {
                // Remove the unwanted param from the URL
                url_to_clean = clean_url(url_to_clean, unwanted_param)
            } else {
                // Split unwanted_param into domain name and the actual param
                parts := strings.Split(unwanted_param, "?")
                unwanted_param_domain_name, unwanted_param_without_domain_name := parts[0], parts[1]
                // Strip the protocol from the URL
                url_without_protocol := regexp.MustCompile("https?://").ReplaceAllString(url_to_clean, "")
                // Strip everything after the domain name from the URL
                url_to_clean_domain_name := regexp.MustCompile("/.*$").ReplaceAllString(url_without_protocol, "")
                // if url_to_clean's domain name contains param's domain name
                if strings.Contains(url_to_clean_domain_name, unwanted_param_domain_name) {
                    // Remove the unwanted param from the URL
                    url_to_clean = clean_url(url_to_clean, unwanted_param_without_domain_name)
                }
            }
        }
    // Print the cleaned URL
    fmt.Println(url_to_clean)
    }
}
