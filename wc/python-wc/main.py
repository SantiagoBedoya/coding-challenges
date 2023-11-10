import argparse
import sys
from counter import count_file_bytes, count_file_characters, count_file_lines, count_file_words

parser = argparse.ArgumentParser()

# arguments
parser.add_argument("file", help="file to read", nargs="?", type=argparse.FileType('rb'), default=sys.stdin)
parser.add_argument("-c", action="store_true", help="count file bytes")
parser.add_argument("-l", action="store_true", help="count file lines")
parser.add_argument("-w", action="store_true", help="count file words")
parser.add_argument("-m", action="store_true", help="count file characters")


args = parser.parse_args()
if args.file == sys.stdin:
    content = sys.stdin.buffer.read()
else:
    content = args.file.read()


values = []
# count file bytes
if args.c:
    bytes = count_file_bytes(content)
    values.append(bytes)

# count file lines
if args.l:
    lines = count_file_lines(content)
    values.append(lines)

# count file words
if args.w:
    num_words = count_file_words(content)
    values.append(num_words)

# count file characters
if args.m:
    characters = count_file_characters(content)
    values.append(characters)

# if no use flags
if not args.c and not args.l and not args.w and not args.m:
    bytes = count_file_bytes(content)
    lines = count_file_lines(content)
    num_words = count_file_words(content)
    characters = count_file_characters(content)
    values = [bytes, lines, num_words, characters]

if args.file is not None and args.file.name != '<stdin>':
    values.append(args.file.name)

print('\t'.join(map(str, values)))
