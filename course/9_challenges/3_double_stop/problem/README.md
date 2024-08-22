Imagine that `Stop()` is called while a critical error happens. Since the handling of the critical error will also make a call to `Stop()`, `Stop()` will close the stop channel twice, which causes a panic.