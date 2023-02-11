import html
import os
import re
import sys
import urllib.parse


def url_unescape(url_to_unescape=None) -> str:
    if url_to_unescape is None or url_to_unescape.strip() == '':
        raise ValueError(
            'got an empty value for the required param url_to_unescape')
        sys.exit(1)
    url_to_unescape = url_to_unescape.strip('\n')
    regex_patterns = {
        'html_entity': r'[&]([a-z]+)|([#][0-9]+)[;]',
        'escaped': r'%[uU]([0-9A-Fa-f]{4})',
        'percent_encoded': r'%([0-9A-Fa-f]{2})'
    }
    how_many_passes = 3
    for _ in range(how_many_passes):
        if re.search(regex_patterns['html_entity'], url_to_unescape):
            url_to_unescape = html.unescape(url_to_unescape)
        if re.search(regex_patterns['escaped'], url_to_unescape):
            url_to_unescape = re.sub(regex_patterns['escaped'], lambda m: chr(
                int(m.group(1), 16)), url_to_unescape)
        if re.search(regex_patterns['percent_encoded'], url_to_unescape):
            url_to_unescape = urllib.parse.unquote(url_to_unescape)
    return url_to_unescape


if __name__ == '__main__':
    print(
        url_unescape(sys.argv[1])
    )
