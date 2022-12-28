import os
import urllib.parse
url_to_unescape = os.getenv('url_to_unescape')
url_to_unescape = url_to_unescape.strip('\n')
unescaped_url = urllib.parse.unquote(url_to_unescape)
print(f'\n{unescaped_url}\n')