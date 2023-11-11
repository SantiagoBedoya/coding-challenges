import sys


def failure(message="invalid JSON syntax"):
    print(message)
    sys.exit(1)


def success():
    sys.exit(0)


def checker(content: str):
    content = content.replace("\n", "")
    if len(content) == 0:
        print("here1")
        failure()

    if content[0] != "{":
        print("here2")
        failure()

    if content[len(content) - 1] != "}":
        failure()

    key_pairs = content.removeprefix("{").removesuffix("}").split(",")
    
    for key_pair in key_pairs:
        if len(key_pair) == 0:
            failure()

        key, value = key_pair.split(":", 1)
        check_key(key.strip())
        check_value(value.strip())


def check_key(key: str):
    if key.count('"') != 2:
        failure()


def check_value(value: str):
    if "{" in value:
        if "}" not in value:
            failure()
        else:
            if len(value) == 2:
                return
            
            k, v = value.removeprefix('{').removesuffix('}').strip().split(':')
            check_key(k.strip())
            check_value(v.strip())

    if "[" in value:
        if "]" not in value:
            failure('invalid boxes')
        else:
            values = value.removeprefix('[').removesuffix(']').strip().split(',')
            for v in values:
                if len(v) == 0:
                    return
                check_value(v)
            return
    
    if '"' in value:
        if value.count('"') % 2 != 0:
            failure('invalid quotes')
        else:
            return


    if value not in ["true", "false", "null"]:
        try:
            int(value)
        except ValueError as e:
            failure('invalid number')
