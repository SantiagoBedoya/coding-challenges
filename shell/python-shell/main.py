import os
import atexit
import readline

histfile = os.path.join(os.path.expanduser('~'), '.python-shell_history')

try:
    # load history file
    readline.read_history_file(histfile)
    readline.set_history_length(1000)
    while True:
        try:
            command = input("python-shell> ").strip()
            if command == "exit":
                break
            try:
                if "cd" in command:
                    os.chdir("/".join([os.getcwd(), command.split(" ")[1]]))
                elif command == "history":
                    with open(histfile) as file:
                        print(''.join(file.readlines()))
                else:
                    os.system(command)
            except FileNotFoundError as e:
                print(e)
            finally:
                atexit.register(readline.write_history_file, histfile)
        except (KeyboardInterrupt, EOFError) as e:
            break
except FileNotFoundError:
    pass





