import sys

if __name__ == '__main__':
    from src import bolinette
    bolinette.run_command(*sys.argv[1:])