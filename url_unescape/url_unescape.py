import argparse
import html
import logging
import re
import sys
import urllib.parse
from dot_dict.dot_dict import DotDict
from journal_logger.journal_logger import JournalLogger


logging.basicConfig(level=logging.DEBUG)
jl = JournalLogger(program_name='url_unescape')


def url_unescape(url_to_unescape=None) -> str:
    if url_to_unescape is None or url_to_unescape.strip() == '':
        jl.print('got an empty value for the required param url_to_unescape')
        return ''
    url_to_unescape = url_to_unescape.strip('\n')
    tasks = DotDict({
        'html_entity': {'regex': r'&[a-z]+|(#[0-9]+);', },
        'escaped': {'regex': r'%[uU]([0-9A-Fa-f]{4})', },
        'percent_encoded': {'regex': r'%[0-9A-Fa-f]{2}', }
    })
    jl.print(f'url_to_unescape: {url_to_unescape!r}')
    while True:
        tasks_performed = []
        if re.search(tasks.html_entity.regex, url_to_unescape):
            result = html.unescape(url_to_unescape)
            is_same = ' (same)'
            if result != url_to_unescape:
                url_to_unescape = result
                is_same = ''
                tasks_performed.append('html_entity')
            jl.print(f'apply html.unescape(): result{is_same}: {result}')
        if re.search(tasks.escaped.regex, url_to_unescape):
            result = re.sub(tasks.escaped.regex,
                            lambda m: chr(int(m.group(1), 16)),
                            url_to_unescape)
            is_same = ' (same)'
            if result != url_to_unescape:
                url_to_unescape = result
                is_same = ''
                tasks_performed.append('escaped')
            jl.print(f'resolve escaped characters (%uXXXX): result{is_same}: {result}')
        if re.search(tasks.percent_encoded.regex, url_to_unescape):
            result = urllib.parse.unquote(url_to_unescape)
            is_same = ' (same)'
            if result != url_to_unescape:
                url_to_unescape = result
                is_same = ''
                tasks_performed.append('percent_encoded')
            jl.print(f'resolve percent-encoded characters (%XX): result{is_same}: {result}')
        if not tasks_performed:
            break
    return url_to_unescape


if __name__ == '__main__':
    usage = """Unescape an escaped URL or any string by running

    python3 url_unescape.py "url or string to unescape"

The script makes a URL or any string prettier by resolving HTML entities (e.g. `&amp;` -> `&`) with Python's `html.unescape`, non-ASCII escaped characters (%%uXXXX) with a regex pattern, and percent-encoded characters (%%XX) with `urllib.parse.unquote`.

Source code: https://github.com/kirisakow/url_tools
----------------------------------------------------------------
"""
    parser = argparse.ArgumentParser(
        usage=usage, epilog="Source code: https://github.com/kirisakow/url_tools")
    parser.add_argument('url_to_unescape', default=None, type=str,
                        help="(required) a string (most often a URL) you want to unescape.")
    args = parser.parse_args()
    try:
        unescaped_url = url_unescape(args.url_to_unescape)
        print(unescaped_url)
    except KeyboardInterrupt:
        sys.exit('\ninterrupted by user')
