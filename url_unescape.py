import html
import os
import urllib.parse
url_to_unescape = os.getenv('url_to_unescape')
url_to_unescape = url_to_unescape.strip('\n')
unescaped_url = html.unescape(url_to_unescape)
unescaped_url = urllib.parse.unquote(unescaped_url)
print(f'\n{unescaped_url}\n')
