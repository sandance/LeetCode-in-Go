import datetime
import json
import threading
import urllib.request
import sys

def main(args):
    if len(args) != 1:
        print('Expecting exactly one argument, got: {!r}'.format(args))
        sys.exit(1)
    mode = args[0]

    # httpbin.org provides several convenient test endpoints. We're using
    # httpbin.org/delay/N, which just waits N seconds before responding.
    urls = [
        'https://httpbin.org/delay/1',
        'https://httpbin.org/delay/2',
        'https://httpbin.org/delay/1',
    ]
    def log_responses(responses):
        log('Got responses: {}'.format(json.dumps(responses, indent=4)))

    if mode == 'parallel':
        start_http_get_parallel(urls, log_responses)
    elif mode == 'serial':
        start_http_get_serial(urls, log_responses)
    else:
        print('Expecting "serial" or "parallel", got {}.'.format(q(mode)))
        sys.exit(1)

    log('Started!')

def log(message):
    ts = datetime.datetime.now().strftime('%H:%M:%S.%f')[:-5]
    print('[{}] {}'.format(ts, message))

def q(s):
    return json.dumps(s)

if __name__ == '__main__':
    main(sys.argv[1:])
