import shutil, os, re

numPattern = re.compile(r"^(\d+)(_.*)$")

if __name__ == "__main__":
    for file in os.listdir():
        res = numPattern.search(file)
        if res :
            length = len(res.group(1))
            left = 4 - length
            print("0"*left + res.group(1)+ res.group(2))
            shutil.move(file, "0"*left + res.group(1)+ res.group(2))

    os.system('pause')