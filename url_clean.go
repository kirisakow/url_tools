package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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

func get_abs_path(filename string) string {
	// Get the path of the executable
	exe_path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Extract the directory portion of the path
	exe_dir := filepath.Dir(exe_path)
	// return absolute path to filename
	return filepath.Join(exe_dir, filename)
}

func unwanted_query_params(filename string) <-chan string {
	ch := make(chan string)
	go func() {
		file, err := os.Open(get_abs_path(filename))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			ch <- strings.TrimSpace(line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		close(ch)
	}()
	return ch
}

// Removes the unwanted query parameter from the URL.
func remove_param_if_present(url_to_clean, unwanted_qparam string) string {
	// If the unwanted param contains an asterisk, replace it with a regex pattern
	unwanted_qparam = strings.ReplaceAll(unwanted_qparam, "*", "[^&=?]*")
	// Use a regular expression to match the unwanted query parameter and remove it
	r := regexp.MustCompile(fmt.Sprintf(`[&]?%s=[^&]*`, unwanted_qparam))
	url_to_clean = r.ReplaceAllString(url_to_clean, "")
	// remove final ? symbol if no params left
	url_to_clean = regexp.MustCompile("[?]$").ReplaceAllString(url_to_clean, "")
	return url_to_clean
}

func clean_url_from_unwanted_query_params(url_to_clean string) string {
	// Iterate over the unwanted query params file contents line by line, as a channel
	for unwanted_qparam := range unwanted_query_params("unwanted_query_params.txt") {
		// Check if the unwanted param does not contain a '?' symbol
		if !strings.Contains(unwanted_qparam, "?") {
			// Remove the unwanted param from the URL
			url_to_clean = remove_param_if_present(url_to_clean, unwanted_qparam)
		} else {
			// Split unwanted_qparam into domain name and the actual param
			parts := strings.Split(unwanted_qparam, "?")
			unwanted_qparam_domain_name, unwanted_qparam_without_domain_name := parts[0], parts[1]
			// Strip the protocol from the URL
			url_without_protocol := regexp.MustCompile("https?://").ReplaceAllString(url_to_clean, "")
			// Strip everything after the domain name from the URL
			url_to_clean_domain_name := regexp.MustCompile("/.*$").ReplaceAllString(url_without_protocol, "")
			// if url_to_clean's domain name contains param's domain name
			if strings.Contains(url_to_clean_domain_name, unwanted_qparam_domain_name) {
				// Remove the unwanted param from the URL
				url_to_clean = remove_param_if_present(url_to_clean, unwanted_qparam_without_domain_name)
			}
		}
	}
	return url_to_clean
}

func main() {
	// Read input either from stdin or from the argument, if any
	urls_to_clean := read_input()
	// Go through the list of URLs to clean
	for _, url_to_clean := range urls_to_clean {
		clean_url := clean_url_from_unwanted_query_params(url_to_clean)
		// Print the cleaned URL
		fmt.Println(clean_url)
	}
}
