from ctypes import cdll

cdll.LoadLibrary("./libadd.so").add()
