package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	outCh := make(Bi)

	for _, stage := range stages {
		in = stage(in)
	}

	go func() {
		defer close(outCh)

		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				if v != nil {
					outCh <- v
				}
			}
		}
	}()

	return outCh
}
