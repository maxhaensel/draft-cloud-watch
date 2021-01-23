from ctypes import *

# loading shared object
lib = cdll.LoadLibrary("main.so")


# go type
# class GoSlice(Structure):
#     _fields_ =


# class Foo(Structure):
#     _fields_ = [("a", c_int), ("b", c_int), ("c", c_int), ("d", c_int), ("e", c_int), ("f", c_int)]

libyy = lib.Foo

libyy.argtypes = [("data", POINTER(c_void_p)), ("len", c_longlong), ("cap", c_longlong)]
# lib.Foo.restype = Foo

lenItem = 6

# t = GoSlice((c_void_p * lenItem)(5, 2, 3, 4, 5, 6), lenItem)
t = GoSlice((c_void_p)(), 0)
f = lib.Foo()

# print(f)
# print(f.a)
# print(f.b)
# print(f.c)
# print(f.d)
# print(f.e)
# print(f.f)
