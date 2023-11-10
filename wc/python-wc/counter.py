def count_file_bytes(content:bytes) -> int:
    return len(content)

def count_file_characters(content:bytes) -> int:
    content_str = content.decode()
    return len(content_str)

def count_file_words(content:bytes) -> int:
    num_words = len(content.split())
    return num_words
    
def count_file_lines(content:bytes) -> int:
    lines = len(content.splitlines())
    return lines
