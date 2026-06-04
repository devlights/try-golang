import ctypes

def main():
    lib = ctypes.CDLL('./libgoadd.so')

    x = 111
    y = 222
    z = lib.GoAdd(x, y)
    
    print(f'[FROM PYTHON] {z}')

if __name__ == '__main__':
    main()
