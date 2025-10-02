# Interfaces using http

Using io.Copy you can copy information from a source to a destination until either end of file is reached or an error occurs. Returns number of bytes copied and the first error encounted while copying, if any.

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


