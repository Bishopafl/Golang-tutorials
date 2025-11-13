# Reading files withing our local machine

Using os.Open opens the named file for reading. If successful, methods on the returned file can be used for reading; the associated file descriptor has mode O_RDONLY. If there is an error, it will be of type *PathError

Doing immediate error handling and exiting the os immediately after the execution of printing the error occurs.

Utilizing io.Copy you can copy information from a source to a destination until either end of file is reached or an error occurs. Returns number of bytes copied and the first error encounted while copying, if any. We are doing this to read the contents of any file passed as an arguement for this go application.

Successful copy returns err == nil not err == EOF. Because Copy is defined to read from source until end of file, it does not treat and end of file from Read as an error to be reported.

----

io.Copy -> something that implements writer interface -> something that implements Reader interface
                            |                                               |
                            v                                               v
                        os.Stdout                                       resp.Body
                            |
                            v
                    value of type File
                            |
                            v
            File has a function called 'Write'
                            |
                            v
        Therefore, it implements the 'Write' interface


