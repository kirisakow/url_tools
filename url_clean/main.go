package main

import (
	"fmt"
	"kirisakow/url_tools/url_cleaner"
)

// test 1:
// url_deref "https://click.e.economist.com/?qs=d9790a138437761ebe4a7038a51a8d29745743f90c5f34402ab029e8ab5f241ae778f7b72fd72dd4a4267b66301268278accbc7214a65cb014407b41816111bf" | url_unescape | url_clean
// test 2:
// url_clean "https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else?utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=20230121&utm_content=ed-picks-article-link-6&etear=nl_special_6&utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=1/21/2023&utm_id=1456534"
func main() {
	// Read input either from stdin or from the argument, if any
	urls_to_clean := url_cleaner.Read_input()
	// Go through the list of URLs to clean
	for _, url_to_clean := range urls_to_clean {
		clean_url := url_cleaner.Clean_url_from_unwanted_query_params(url_to_clean)
		// Print the cleaned URL
		fmt.Print(clean_url)
	}
}
