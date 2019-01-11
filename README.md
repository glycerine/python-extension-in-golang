# sum: calling golang library from python

Sample code showing how to embed Go libs in python3:

https://hackernoon.com/extending-python-3-in-go-78f3a69552ac

~~~
(env3) jaten@jatens-MacBook-Pro ~/go/src/github.com/glycerine/sum (master) $ make
go build -buildmode=c-shared -o sum.so sum.go
python build.py > sum.c
gcc sum.c -dynamiclib sum.so -o testsum.so -I/usr/local/Cellar/python/3.7.1/Frameworks/Python.framework/Versions/3.7/include/python3.7m -I/usr/local/Cellar/python/3.7.1/Frameworks/Python.framework/Versions/3.7/include/python3.7m -Wno-unused-result -Wsign-compare -Wunreachable-code -fno-common -dynamic -DNDEBUG -g -fwrapv -O3 -Wall -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk/System/Library/Frameworks/Tk.framework/Versions/8.5/Headers -L/usr/local/opt/python/Frameworks/Python.framework/Versions/3.7/lib/python3.7/config-3.7m-darwin -lpython3.7m -ldl -framework CoreFoundation
go build -buildmode=c-shared -o better.so better.go
python
python build2.py > better.c
gcc better.c -dynamiclib better.so -o testbetter.so -I/usr/local/Cellar/python/3.7.1/Frameworks/Python.framework/Versions/3.7/include/python3.7m -I/usr/local/Cellar/python/3.7.1/Frameworks/Python.framework/Versions/3.7/include/python3.7m -Wno-unused-result -Wsign-compare -Wunreachable-code -fno-common -dynamic -DNDEBUG -g -fwrapv -O3 -Wall -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk/System/Library/Frameworks/Tk.framework/Versions/8.5/Headers -L/usr/local/opt/python/Frameworks/Python.framework/Versions/3.7/lib/python3.7/config-3.7m-darwin -lpython3.7m -ldl -framework CoreFoundation
# do we need to add -fPIC -shared ?

(env3) jaten@jatens-MacBook-Pro ~/go/src/github.com/glycerine/sum (master) $ python
Python 3.7.1 (default, Nov  6 2018, 18:45:35) 
[Clang 10.0.0 (clang-1000.11.45.5)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> import testbetter
>>> testbetter
<module 'testbetter' from '/Users/jaten/go/src/github.com/glycerine/sum/testbetter.so'>
>>> dir(testbetter)
['NewDictionary', 'Sum', 'SumOverSlice', '__doc__', '__file__', '__loader__', '__name__', '__package__', '__spec__']
>>> testbetter.SumOverSlice([1,5,7])
13
>>> testbetter.Sum(4,8)
12
>>> 
~~~
