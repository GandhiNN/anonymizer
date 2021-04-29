package hasher

import (
	"bytes"
	"io"
	"os/exec"
)

// SimplePipedCommand : emulate `echo $MSISDN | md5sum`
func SimplePipedCommand(cmd1, cmd2, msisdn string) *bytes.Buffer {

	// create commands
	c1 := exec.Command(cmd1, msisdn)
	c2 := exec.Command(cmd2)

	// make a pipe and bytes buffer
	reader, writer := io.Pipe()
	var buf bytes.Buffer

	// set the output of c1 to pipe writer
	c1.Stdout = writer

	// set the input of c2 to pipe reader
	c2.Stdin = reader

	// cache the output of c2 to memory
	c2.Stdout = &buf

	// execute c1 and c2
	c1.Start()
	c2.Start()

	// wait for c1 to complete and close the writer
	c1.Wait()
	writer.Close()

	// wait for c2 to complete and close the reader
	c2.Wait()
	reader.Close()

	return &buf
}

// BashPipedCommand : for multiple piped commands invocation
//  i.e. `echo a | md5sum | awk '{print $1}'`
func BashPipedCommand(cmds ...*exec.Cmd) (pipeOut, stdErrs []byte, pipeErr error) {

	if len(cmds) < 1 {
		return nil, nil, nil
	}

	// memory container to collect the output and stderr from the command(s)
	var output bytes.Buffer
	var stderr bytes.Buffer

	last := (len(cmds) - 1)
	for i, cmd := range cmds[:last] {
		var err error
		// stdout of prev command as input to stdin of the executed command
		cmds[i+1].Stdin, err = cmd.StdoutPipe()
		if err != nil {
			return nil, nil, err
		}
		// collect each command's stderr
		cmd.Stderr = &stderr
	}

	// we do not collect all output
	//  just collect the stdout and stderr of the last command
	cmds[last].Stdout, cmds[last].Stderr = &output, &stderr

	// start each command
	for _, cmd := range cmds {
		err := cmd.Start()
		if err != nil {
			return output.Bytes(), stderr.Bytes(), err
		}
	}

	// wait for each command to complete
	for _, cmd := range cmds {
		err := cmd.Wait()
		if err != nil {
			return output.Bytes(), stderr.Bytes(), err
		}
	}

	// return the pipeline output and the collected standard error
	return output.Bytes(), stderr.Bytes(), nil
}
