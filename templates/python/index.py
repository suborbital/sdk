"""File run by SE2"""

import sdk

import lib


def run_e(payload, ident):
    """Function for SE2"""
    lib.log(f'Received payload "{payload}"', ident)
    result = lib.run(payload)

    lib.log(f'Returning result "{result}"', ident)
    sdk.return_result(result, ident)

    lib.log(f'Result return for "{ident}"', ident)
