import html
import os
import logging
import re
import sys
import urllib.parse
sys.path.append(
    os.path.abspath(
        os.path.join(
            os.path.dirname(__file__),
            '../telegram_bots'
        )
    )
)
from utils import JournalLogger

logging.basicConfig(level=logging.DEBUG)
jl = JournalLogger(program_name='url_tools')


def url_unescape(url_to_unescape=None) -> str:
    if url_to_unescape is None or url_to_unescape.strip() == '':
        jl.print('got an empty value for the required param url_to_unescape')
        return ''
    url_to_unescape = url_to_unescape.strip('\n')
    tasks = {
        'html_entity': {'regex': r'&[a-z]+|(#[0-9]+);', 'repeat': True},
        'escaped': {'regex': r'%[uU]([0-9A-Fa-f]{4})', 'repeat': True},
        'percent_encoded': {'regex': r'%([0-9A-Fa-f]{2})', 'repeat': True}
    }
    jl.print(f'url_to_unescape: {url_to_unescape}')
    while all([task['repeat'] == True for task in tasks.values()]):
        if (tasks['html_entity']['repeat'] == True
                and re.search(tasks['html_entity']['regex'], url_to_unescape)):
            jl.print(f'apply html.unescape()')
            result = html.unescape(url_to_unescape)
            jl.print(f'result: {result}')
            if result != url_to_unescape:
                url_to_unescape = result
            else:
                tasks['html_entity']['repeat'] = False
        if (tasks['escaped']['repeat'] == True
                and re.search(tasks['escaped']['regex'], url_to_unescape)):
            jl.print(f"resolve escaped characters (%uXXXX)")
            result = re.sub(tasks['escaped']['regex'],
                            lambda m: chr(int(m.group(1), 16)),
                            url_to_unescape)
            jl.print(f'result: {result}')
            if result != url_to_unescape:
                url_to_unescape = result
            else:
                tasks['escaped']['repeat'] = False
        if (tasks['percent_encoded']['repeat'] == True
                and re.search(tasks['percent_encoded']['regex'], url_to_unescape)):
            jl.print(f"resolve percent-encoded characters (%XX)")
            result = urllib.parse.unquote(url_to_unescape)
            jl.print(f'result: {result}')
            if result != url_to_unescape:
                url_to_unescape = result
            else:
                tasks['percent_encoded']['repeat'] = False
    return url_to_unescape


if __name__ == '__main__':
    try:
        unescaped_url = url_unescape(args.url_to_unescape)
        print(unescaped_url)
    except KeyboardInterrupt:
        sys.exit('\ninterrupted by user')
