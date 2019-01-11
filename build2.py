from pybindgen import retval, param, Module
import sys

mod = Module('testbetter')
mod.add_include('"better.h"')
mod.add_function('Sum', retval('int'), [param('int', 'a'), param('int', 'b')])
mod.add_function('NewDictionary', retval('PyObject *', caller_owns_return=False), [])

mod.add_function('SumOverSlice', retval('int'), [param('PyObject *', 'slice', transfer_ownership=True)])

mod.generate(sys.stdout)