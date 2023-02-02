package url_cleaner

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//go:embed unwanted_query_params.txt
var embedded_assets embed.FS

func get_regex_pattern(key string) string {
	regex_pattern := make(map[string]string)
	regex_pattern["http_protocol"] = `https?://`
	regex_pattern["http_url"] = regex_pattern["http_protocol"] + `[a-zA-Z0-9_.](:[0-9]{2,5})?(\S+)?`
	if val, key_exists := regex_pattern[key]; key_exists {
		return val
	}
	panic(fmt.Sprintf("unknown key %q", key))
}

// Reads input either from stdin or from the argument, if any.
// Returns the input as a string.
func Read_input() []string {
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

func search_recursively_for_filename(root_dir, filename string) (string, error) {
	var abs_location string
	// Search recursively for filename
	search_err := filepath.Walk(root_dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == filename {
			abs_location = path
			return nil
		}
		return nil
	})
	return abs_location, search_err
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
	// Search recursively for filename
	abs_location, err := search_recursively_for_filename(exe_dir, filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// return absolute path to filename
	return abs_location
	// return filepath.Join(exe_dir, filename)
}

func unwanted_query_params(filename string) <-chan string {
	ch := make(chan string)
	go func() {
		abs_path_to_unw_qp := get_abs_path(filename)
		if abs_path_to_unw_qp != "" {
			// Read unwanted query params file from disk
			if unw_qp_file, err := os.Open(abs_path_to_unw_qp); err == nil {
				defer unw_qp_file.Close()
				scanner := bufio.NewScanner(unw_qp_file)
				for scanner.Scan() {
					line := scanner.Text()
					ch <- strings.TrimSpace(line)
				}
				if err := scanner.Err(); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		} else {
			// Read unwanted query params file from the embedded filesystem
			log.Printf("error while reading %q from disk\n", filename)
			log.Printf("trying to retrieve %q from an embedded filesystem with the Open() method\n", filename)
			if unw_qp_file_embedded, err := embedded_assets.Open(filename); err == nil {
				scanner := bufio.NewScanner(unw_qp_file_embedded)
				for scanner.Scan() {
					line := scanner.Text()
					ch <- strings.TrimSpace(line)
				}
				if err := scanner.Err(); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			} else {
				log.Printf("error while reading %q from an embedded filesystem with the Open() method: %s", filename, err)
				log.Printf("trying to retrieve %q from an embedded filesystem with the ReadFile() method\n", filename)
				if unw_qp_file_embedded_as_bytearray, err := embedded_assets.ReadFile(filename); err == nil {
					lines := strings.Split(string(unw_qp_file_embedded_as_bytearray), "\n")
					for _, line := range lines {
						ch <- strings.TrimSpace(line)
					}
				}
			}
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
	// remove any of the symbols ?&# or any combination of them if found at the end of the URL string
	url_to_clean = regexp.MustCompile("[?&#]+$").ReplaceAllString(url_to_clean, "")
	return url_to_clean
}

func Clean_url_from_unwanted_query_params(url_to_clean string) string {
	is_valid_url := regexp.MustCompile(get_regex_pattern("http_url")).MatchString(url_to_clean)
	if !is_valid_url {
		return fmt.Sprintf("input string %q is not a valid URL", url_to_clean)
	}
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
			url_without_protocol := regexp.MustCompile(get_regex_pattern("http_protocol")).ReplaceAllString(url_to_clean, "")
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
