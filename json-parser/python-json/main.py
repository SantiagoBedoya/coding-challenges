import argparse
from json_parser import checker

parser = argparse.ArgumentParser()

# arguments
parser.add_argument("file", help="json file to parse", type=argparse.FileType('r'))

# parse arguments
args = parser.parse_args()

# read content from file
content = args.file.read()

# check the content
response = checker(content)