"""File to be edited by users"""


def run(payload_string):
    """ 'run' is imported by index.py"""
    reversed_string = ""
    for i in payload_string:
        reversed_string = i + reversed_string
    return reversed_string


def log(msg, msg_id=None):
    """Log message id"""
    print(f'\033[35m [plugin.py]\033[0m | id={msg_id} | {msg}', flush=True)