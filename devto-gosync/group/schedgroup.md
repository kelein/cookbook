---
marp: true
---

schedgroup: a timer-based goroutine concurrency primitive
16 Jun 2020

Matt Layher
Distributed Systems Engineer, Fastly
mdlayher@gmail.com
https://github.com/mdlayher
https://mdlayher.com
@mdlayher

---

: Hi everyone, I'm Matt Layher. Today I'm going to share the lessons I've learned
: about Go concurrency and timers while building out a package I call schedgroup;
: short for "scheduler group".

* Matt Layher

@mdlayher, github.com/mdlayher/talks, twitch.tv/mdlayher

---

: You can find me on GitHub and Twitter at mdlayher, and these slides and code
: examples will be available at my talks repository on GitHub.
: ---
: I've also recently started streaming Go and network programming on Twitch, so
: please stop by and join the chat!

.image colorado-square.jpg 500 _

* Introducing schedgroup

A package that allows scheduling or delaying a goroutine to run at or after a
specific point in time.

.play basic/main.go /START OMIT/,/END OMIT/

: The schedgroup package allows delaying the execution of a goroutine until after
: a specific point in time.
: ---
: The New constructor creates a Group and allows cancelation via context.
: ---
: Next, the inner loop will schedule goroutines that print a number to the screen
: every 500 milliseconds.
: ---
: Finally, the Wait method will wait for all of the scheduled tasks to complete.
: ---
: (run example)

.link https://github.com/mdlayher/schedgroup

* Why?

Typically you want your code to run immediately and as fast as possible.

Sometimes you don't make the rules!

: So why build out a package that will only make your code slower? Sometimes you
: have to follow someone else's rules.

* Introducing CoreRAD

CoreRAD is my take on a modern IPv6 router advertisement daemon.

.link https://github.com/mdlayher/corerad

- Allows IPv6 clients to find routers, establish a default route, and configure
  IPv6 addresses

CoreRAD uses NDP (RFC 4861) to perform these functions:

.link https://tools.ietf.org/html/rfc4861

Learn more about IPv6 and NDP:

.link https://mdlayher.com/talks

: My current side project is CoreRAD, a modern take on an IPv6 router advertisement
: daemon. The schedgroup package originated within CoreRAD before being factored
: out into its own repository.
: ---
: CoreRAD and IPv6 are not the focus of today's talk, but you can learn more by
: checking out my talks page and joining us in #networking on Gophers Slack.

* NDP rules

RFC 4861 mandates that IPv6 routers rate limit their responses.

When multicasting to the IPv6 link-local all nodes address:

    ... consecutive Router Advertisements sent to the all-nodes multicast address MUST be rate
    limited to no more than one advertisement every MIN_DELAY_BETWEEN_RAS seconds.

And even when unicasting a response to individual router solicitations:

    In all cases, Router Advertisements sent in response to a Router Solicitation MUST be
    delayed by a random time between 0 and MAX_RA_DELAY_TIME seconds.

: CoreRAD implements the Neighbor Discovery Protocol, and the NDP RFC requires us
: to delay certain responses for a specified period of time.
: ---
: For multicasts, we can only send a router advertisement every 3 seconds.
: ---
: For unicasts, we must delay each response by a random time between 0 and
: 500 milliseconds.

* time.AfterFunc: using the standard library

`time.AfterFunc` does exactly what we need! Problem solved?

.play naive/main.go /START OMIT/,/END OMIT/

: As a first pass, let's implement this logic using time.After.
: ---
: We create a WaitGroup and prepare it to synchronize 5 goroutines. The goroutines
: will print a number to the screen every 500 milliseconds. Once all of the
: goroutines are finished, we print "done" to the screen.
: ---
: (run example)
: ---
: This demo behaves identically to the schedgroup demo from earlier.

* Cancelation support: the naive prototype

Suppose we want to schedule many tasks and also offer context cancelation.

.play naivecontext/main.go /START OMIT/,/END OMIT/

: Now that we have a prototype, what would it take to add context cancelation
: support?
: ---
: Suppose we want to schedule a million tasks to atomically increment an integer,
: but also halt the work early once all of the goroutines have been created.
: ---
: (run example)
: ---
: We can see that some portion of the tasks did complete before they were canceled.
: But what are most of those goroutines doing while they wait for 1 second to elapse?

* Cancelation support: the naive prototype

We can use pprof to examine our program with a nice web UI:

    $ go tool pprof -http :8080 localhost:8081/debug/pprof/goroutine

: pprof is one of the best Go debugging and performance analysis tools. This command
: will produce a nice graph of goroutine activity in our program.

* Cancelation support: the naive prototype

.image naivecontext/goroutines.png 550 _

: As we can see, the majority of the goroutines are parked and doing absolutely
: nothing other than consuming memory while our program is running.

* Cancelation support: the naive prototype

No library, the caller handles goroutines and timers.

Pros:

- Simplicity: very little code

Cons:

- Caller has to manage concurrency and cancelation
- Unbounded number of goroutines, wasteful use of memory
- Excessive number of runtime timers

: These prototypes are small and easy to comprehend, but I think we can improve
: on the idea by limiting the number of goroutines and timers which will immediately
: block on a task.
: ---
: It's worth noting that Go 1.14 was released with a significant rework to the runtime
: timer code. I started building schedgroup during the 1.13 cycle, so perhaps some of
: my concerns about timer efficiency are less of an issue with 1.14.

* Efficiency vs. convenience

Can we write efficient and tidy code which also supports cancelation?

- `time.AfterFunc` is well-optimized, but inflexible
- Spinning up a goroutine to wait for a timer or cancelation is wasteful

: We've hit a classic programming trade-off: how do we balance efficiency vs.
: convenience?
: ---
: Can we have the best of both worlds by only spinning up goroutines as needed,
: while also allowing context cancelation for callers?

* schedgroup: worker pool

A library with a fixed number of worker goroutines consuming work from a channel.

Goals:

- Bounded number of goroutines
- Support for context cancelation

: For the first library iteration, I decided to try a worker pool approach with
: a fixed number of worker goroutines.
: ---
: Our first goal is to provide a concise API which supports context cancelation.

* schedgroup: worker pool

A Group coordinates concurrent workers to expose a concise public API.

.code workerpool/schedgroup/group.go /START 1 OMIT/,/END 1 OMIT/

: The Group is our top-level type, and the New constructor creates a Group while
: also spinning up our goroutine workers.
: ---
: I decided to create 32 goroutines and a buffered channel to enqueue tasks,
: with the intent of providing more flexibility later if needed.

* schedgroup: worker pool

Tasks are created by Delay, and the caller can block until completion by calling Wait.

.code workerpool/schedgroup/group.go /START 2 OMIT/,/END 2 OMIT/

: The exported methods of our API are Delay and Wait.
: ---
: Delay schedules the execution of a function at or after the input delay. I mention
: "at or after" specifically because the scheduling is best effort and may be
: influenced by events such as goroutine scheduling or garbage collection.
: In practice this hasn't been a problem for me or my applications.
: ---
: Wait blocks until all of the worker goroutines complete, allowing external
: synchronization by the caller.

* schedgroup: worker pool

Internal worker goroutines consume tasks and execute them when a delay elapses.

.code workerpool/schedgroup/group.go /START 3 OMIT/,/END 3 OMIT/

: The worker method is used to create a worker which will run and consume work
: until ctx is canceled.
: ---
: work is called by each worker when a task is received. It either waits for
: context cancelation or for the caller-specified delay to elapse. When the delay
: time elapses, the task is run.

* schedgroup: worker pool

Demo appears and behaves identically to the completed version earlier in the slides.

.play workerpool/main.go /START OMIT/,/END OMIT/

: Putting this all together...
: ---
: (run example)
: ---
: We've come up with a demo which behaves identically to our naive version.

* schedgroup: worker pool

.image workerpool/goroutines.png 550 _

: Looking at the pprof goroutine output for this demo, we can see the number of
: goroutines is stable with 32 workers and a few extras for pprof and such.

* schedgroup: worker pool

A library with a fixed number of worker goroutines consuming work from a channel.

Pros:

- Concurrency and cancelation are managed internally
- Bounded number of goroutines
- Reasonably concise code thanks to Go's concurrency primitives

Cons:

- Idle worker goroutines consume system resources
- An overwhelming amount of tasks could result in blocking the caller

: Overall this seemed like a nice starting point for the schedgroup package.
: The internal code is concise and clean, we support context cancelation, and
: there is no unbounded use of goroutines.
: ---
: However, there are a couple of cons as well. Although it isn't much, the idle
: worker goroutines do consume some system resources. Because of the fixed number
: of workers, it's possible to block the caller when calling Delay if the internal
: buffered channel is already full of tasks.
: ---
: This is an excellent start, but I decided to keep iterating to see what I could
: come up with.

* schedgroup: timer-polling using a monitor goroutine

A library which polls continuously to check for timer expiration.

Goals:

- Don't spin up goroutines until tasks are ready
- Support for context cancelation
- Efficient scheduling of tasks which are ready

: I started asking some folks in the #performance channel on Gophers Slack for
: advice, and Egon Elbre came up with a couple of excellent ideas that heavily
: influenced the final design of the package.
: ---
: First we decided to investigate a timer-polling approach, with the goal of
: having a central monitor goroutine that handles timers and goroutine creation.

* Introducing container/heap

A min-heap will allow us to quickly determine which tasks should be scheduled first.

    // A type, typically a collection, that satisfies sort.Interface can be sorted
    // by the routines in container/heap.
    type Interface interface {
        // sort.Interface
        Len() int
        Less(i, j int) bool
        Swap(i, j int)

        Push(x interface{}) // add x as element Len()
        Pop() interface{}   // remove and return element Len() - 1.
    }

Let's implement `heap.Interface` for a slice of tasks.

: In order to determine which tasks should be scheduled first, we make use of
: a min-heap with Go's container/heap package.
: ---
: To make use of a heap, we have to create a type which implements sort.Interface
: and a couple of extra methods to push items onto and pop them off the heap.

* Implementing heap.Interface

.code timerpolling/schedgroup/group.go /START TASK OMIT/,/END TASK OMIT/

: We wrap our task type in a slice type called tasks, and use that to implement
: heap.Interface. If you've ever implemented sort.Interface, several of the
: methods are identical.
: ---
: Tasks which are scheduled earlier are considered "lesser" in the min-heap.
: Push and Pop are used to add and remove a task to and from the slice respectively.

* container/heap usage

`container/heap` functions maintain the min-heap's structure.

.play timerpolling/heap.go /START HEAP OMIT/,/END HEAP OMIT/

: We can demonstrate that our heap.Interface implementation is correct by creating
: several tasks with different timestamps and pushing them onto the heap in arbitrary order.
: ---
: (run example)
: ---
: Running the example, we see that tasks which are scheduled for sooner will
: be invoked first, causing the numbers to print in increasing order until the
: heap is completely empty.

* schedgroup: timer-polling using a monitor goroutine

A Group coordinates concurrent workers to expose a concise public API.

.code timerpolling/schedgroup/group.go /START GROUP OMIT/,/END GROUP OMIT/

: The Group type must now manage an internal monitor goroutine, along
: with the necessary context to cancel that goroutine later on.

* schedgroup: timer-polling using a monitor goroutine

The Delay method makes a return, but is now a wrapper for the generalized Schedule.

.code timerpolling/schedgroup/group.go /START DELAY OMIT/,/END DELAY OMIT/

: The Delay method also returns, but this time around it becomes a wrapper for the
: generalized Schedule method. Schedule deals with absolute time.Time values
: rather than time.Duration values.
: ---
: As a special case, a negative Delay or past Schedule time will result in a
: task being executed immediately.

* schedgroup: timer-polling using a monitor goroutine

The monitor goroutine is responsible for ticking every millisecond to see if
the tasks heap has any scheduled tasks ready.

.code timerpolling/schedgroup/group.go /START MONITOR OMIT/,/END MONITOR OMIT/

: The monitor goroutine manages the central timer which polls for task readiness.
: I decided to go with the value of 1 millisecond just to see what would happen,
: but this caused some performance problems that we'll dive into later on.

* schedgroup: timer-polling using a monitor goroutine

The trigger method will see if any work is ready to be invoked as of time `now`.

.code timerpolling/schedgroup/group.go /START TRIGGER OMIT/,/END TRIGGER OMIT/

: trigger is invoked regularly by monitor to examine the first task on the heap
: and determine if it is ready to run.
: ---
: Whenever trigger is called, it is possible for more than one task to be ready
: to run, so we will continue the loop until the first task in the heap is no
: longer ready.

* schedgroup: timer-polling using a monitor goroutine

Wait blocks until work is done or the Group's context is canceled.

: Finally, Wait allows the caller to determine if all of the outstanding tasks have
: completed.
: ---
: The Context passed to New can short-circuit the Wait operation, but otherwise
: we continuously poll for readiness until no more tasks are left in the heap.
: ---
: When Wait returns, we cancel the monitor goroutine as well.

.code timerpolling/schedgroup/group.go /START WAIT OMIT/,/END WAIT OMIT/

* schedgroup: timer-polling using a monitor goroutine

Demo appears and behaves identically to both previous versions.

: (run example)
: ---
: Putting it all together, we can see that this version of schedgroup also
: behaves as expected. But there is a big problem with this prototype.

.play timerpolling/main.go /START OMIT/,/END OMIT/

* schedgroup: timer-polling using a monitor goroutine

I deployed this code and noticed a significant CPU usage increase.

What's happening? Let's try out `go` `tool` `trace`.

    $ curl http://localhost:8081/debug/pprof/trace > trace.out
    $ go tool trace trace.out

Note the subtle red/pink banding in the PROCS section.

: When I deployed this code in CoreRAD, I noticed a significant increase in the
: process's CPU usage. To fully understand the situation, I decided to try out
: 'go tool trace' to visualize the process's CPU usage.
: ---
: In the current zoomed-out image, it may be difficult to spot the subtle red/pink
: banding in the PROCS section, but this banding often indicates a problem.

.image timerpolling/trace.png 250 _

* schedgroup: timer-polling using a monitor goroutine

Focusing on a 100ms slice shows the problem more clearly:

- Polling the tasks heap every 1ms results in excessive CPU usage
- Ideally we want sustained CPU usage or no CPU usage; this is just inefficient!

.image timerpolling/tracezoom.png 250 _

: By zooming in on a 100ms slice, we can see repeated, small periods of CPU
: utilization as goroutines thrash between different OS threads running on my CPU.
: ---
: This is a direct effect of the timer-polling schedgroup prototype! For optimal
: efficiency and performance, we want to see sustained CPU usage or no CPU usage at all.

* schedgroup: timer-polling using a monitor goroutine

A library which polls continuously to check for timer expiration.

Pros:

- Concurrency and cancelation are managed internally
- Bounded number of goroutines

Cons:

- Extremely inefficient use of CPU due to wasteful timer polling
- Both Group.monitor and Group.Wait must poll to wait for completion

: This prototype is ultimately a step back from the previous one, but we learned
: some valuable lessons along the way.
: ---
: Be cautious of any sort of polling in your applications, or you may end up
: wasting a lot of CPU time doing effectively no work. Because the monitor and
: Wait methods both have to poll for readiness, this problem is multiplied.
: ---
: We need to take a totally different approach with our next prototype.

* schedgroup: event-driven scheduler with signaling channels

A library which uses goroutines, channels, and select to efficiently signal
when work is ready.

Goals:

- Don't spin up goroutines until tasks are ready
- Support for context cancelation
- Efficient scheduling of tasks which are ready
- No polling of timers to avoid excessive CPU usage

This is the final result of my experiments:

.link https://github.com/mdlayher/schedgroup

: Finally, we have arrived at our final schedgroup design. This prototype
: ultimately became the final package which you can use in your own applications.
: Another conversation with Egon in Gophers Slack also lead to this design.
: ---
: Go's channels which allow passing messages between goroutines in an efficient way.
: The select statement allows us to coordinate sending and receiving channel
: messages to execute different code paths.
: ---
: By combining Go's concurrency primitives and runtime timers, we can produce a
: design that meets all of our needs: on-demand goroutine creation, efficient CPU
: usage, and no timer-polling!

* schedgroup: event-driven scheduler with signaling channels

Some parts of the previous design are retained:

- Min-heap of tasks ordered by their scheduled time
- Monitor goroutine to schedule tasks when they are ready to run

: Some of the good ideas from the previous prototype are carried over into this
: design as well.
: ---
: A min-heap is the perfect data structure for determining which tasks must be
: scheduled first, and the monitor goroutine allows complicated scheduling logic
: to be hidden away from the caller.

* schedgroup: event-driven scheduler with signaling channels

Let's break down the Group architecture:

.image final/diagram.png 400 _

: The Group type has a few exported entry points: Delay/Schedule and Wait. Context
: cancelation is supported throughout to immediately halt scheduling more tasks.
: ---
: The monitor goroutine uses a select statement to wait for context cancelation,
: a new task to be added by schedule, or the a timer tick indicating that
: an existing task is ready to run.
: When monitor invokes trigger, trigger will check if tasks are ready to run.
: If none are, it sends the number of remaining tasks to Wait and returns early.
: ---
: If a task's deadline expires, trigger will execute it immediately by popping it off the
: heap, running it in a goroutine, then reporting the number of remaining tasks to Wait.

* schedgroup: event-driven scheduler with signaling channels

A Group coordinates concurrent workers to expose a concise public API.

: The Group type is still our foundation, but this time we have a couple of internal
: channels which are used for signaling between monitor and calling goroutines.
: ---
: addC sends a notification that a task was added to the heap, and lenC indicates
: the number of tasks left in the heap after a trigger call returns to monitor.

.code final/schedgroup/group.go /START GROUP OMIT/,/END GROUP OMIT/

* schedgroup: event-driven scheduler with signaling channels

Delay and Schedule now signal that work has been added to the heap.

.code final/schedgroup/group.go /START SCHEDULE OMIT/,/END SCHEDULE OMIT/

: Delay and Schedule push a new task onto the heap and then notify the monitor
: goroutine by sending on addC.

* schedgroup: event-driven scheduler with signaling channels

This is a non-blocking send attempt on g.addC:

- If a value can be sent on g.addC, the value is sent.
- Otherwise, the default case is selected and nothing happens.

    // Notify monitor that a new task has been pushed on to the heap.
    select {
    case g.addC <- struct{}{}:
        // empty struct is sent
    default:
        // nothing happens
    }

Why an empty struct? Some prefer using a bool.

I prefer the empty struct because it indicates that the value sent is truly
meaningless and cannot be interpreted in any way.

: Breaking it down further, this is a non-blocking send attempt on addC. If
: monitor is in its select statement waiting for a channel to be ready, the value
: is sent and monitor can proceed. Otherwise, the default case fires and nothing happens.
: Either way, we don't block the caller and the task is scheduled.
: ---
: So why use an empty struct rather than a bool? For me it's a personal preference:
: the empty struct clearly communicates that the value is meaningless and shouldn't
: be perceived as success/failure. Others may feel differently, and that's okay too.

* schedgroup: event-driven scheduler with signaling channels

The monitor goroutine runs a single timer which waits until the earliest
scheduled event is ready to fire.

.code final/schedgroup/group.go /START MONITOR1 OMIT/,/END MONITOR1 OMIT/

: monitor is a bit more complex this time around. After checking for context
: cancelation, it triggers any tasks that are ready as of the current time.
: ---
: trigger now returns the deadline of the next task. If there is a non-zero deadline,
: we reset our timer to fire at that time and initialize tickC for the upcoming select.
: ---
: Otherwise, we stop the timer until another task is scheduled.

* schedgroup: event-driven scheduler with signaling channels

Once trigger returns, we wait until a new event occurs.

.code final/schedgroup/group.go /START MONITOR2 OMIT/,/END MONITOR2 OMIT/

If tickC is nil because there are no more tasks waiting to be scheduled, that
select case can be "shut off".

: This select statement is where our use of Go concurrency primitives truly shines.
: We wait for context cancelation, or addition of a new task, or for tickC to fire,
: indicating another task is ready.
: ---
: If the timer was stopped above because no tasks are in the heap, tickC is nil,
: effectively "shutting off" that case in the select statement. This is a handy trick!

* schedgroup: event-driven scheduler with signaling channels

When trigger returns, we want to notify any outstanding Wait calls if there are
tasks left to be scheduled.

.code final/schedgroup/group.go /START TRIGGER1 OMIT/,/END TRIGGER1 OMIT/

: When trigger is called by monitor, it acquires a mutex to prevent concurrent
: work scheduling. We defer a call which will send the number of items in the heap
: on lenC, signaling Wait. If Wait has been called, it consumes this signal to
: check for completion. If not, the default case fires and nothing happens.

* schedgroup: event-driven scheduler with signaling channels

trigger will run any tasks with a deadline before or equal to now.

.code final/schedgroup/group.go /START TRIGGER2 OMIT/,/END TRIGGER2 OMIT/

: Within the body of trigger, we consume tasks from the heap until there are
: none left, or all remaining tasks must wait for their deadline to pass.
: ---
: Each task which has reached its deadline will be executed in a goroutine to
: avoid blocking trigger.
: ---
: trigger will return either the deadline of the next task in the heap, or a zero
: time. monitor uses these values to determine whether or not to start its timer
: and enable the tickC select case.

* schedgroup: event-driven scheduler with signaling channels

Wait is used to block and wait for outstanding tasks in a Group.

.code final/schedgroup/group.go /START WAIT1 OMIT/,/END WAIT1 OMIT/

: Finally, let's examine Wait. It will block until all tasks in a Group have been
: run, or until the context passed to New is canceled.
: ---
: Both Wait and trigger depend on a mutex to inspect the tasks heap and determine
: if task execution is complete.
: ---
: This code was surprisingly tricky to get right. Typically, defer-unlocking a mutex
: in Go is a best practice, but by using that pattern here, I ran into numerous
: deadlocks in this code! If your program is ever stuck, you can use Ctrl+\ to
: send SIGQUIT to your Go program, immediately dumping the goroutine stacks.

* schedgroup: event-driven scheduler with signaling channels

trigger sends the number of remaining tasks on g.lenC until none are left.

.code final/schedgroup/group.go /START WAIT2 OMIT/,/END WAIT2 OMIT/

: Since trigger sends the number of tasks left in the heap as it returns, we select
: on ctx.Done and lenC to determine when to unblock. Context cancelation is checked
: frequently to ensure we don't do unnecessary work when the caller has asked us
: to abandon the remaining tasks.
: ---
: If no more tasks are left in the heap, we can cancel the monitor goroutine and
: return, thus ending the schedgroup lifecycle.

* schedgroup: event-driven scheduler with signaling channels

Demo appears and behaves identically to all previous versions.

: (run example)
: ---
: The final demo code does exactly what we want, success!

.play final/main.go /START OMIT/,/END OMIT/

* schedgroup: event-driven scheduler with signaling channels

Pros:

- Concurrency and cancelation are managed internally
- Single timer heap management goroutine
- Goroutines are spun up on-demand exactly when needed
- No polling!

Cons:

- Internal complexity: it took me a while to work out bugs
- Potential for deadlocks due to subtle concurrency bugs

: This final event-driven approach is inherently complex and contains some slightly
: tricky uses of Go concurrency, but thanks to a solid test suite, I feel comfortable
: using this code in my production applications.
: ---
: Only a single timer is needed to coordinate tasks, there is no timer polling,
: and goroutines are spun up on demand. In the future, I may add an optional,
: buffered channel semaphore to limit the number of concurrent goroutines executing.

* schedgroup: naive prototype vs. event-driven scheduler

Let's do a head-to-head comparison of our naive `time.AfterFunc` prototype vs. the
final implementation.

We will examine:

- number of running goroutines
- in-use heap space

: Now that we've built out our final prototype, it's time for a head-to-head
: comparison with our initial, naive prototype that uses time.AfterFunc.
: ---
: The CPU usage on both should be comparable due to goroutines sleeping and waiting
: for timers to wake them up, so we'll specifically examine the number of running
: goroutines, and the in-use heap space.

* schedgroup: naive prototype vs. event-driven scheduler

Increment a counter to 1 million after ~10 seconds per task.

.code final/finalpprof/main.go /START OMIT/,/END OMIT/

: For comparison, we will reuse the demo which creates 1 million tasks that atomically
: increment a counter.
: ---
: We create a schedgroup, create 1 million tasks which will sleep for 10 seconds
: each so we can gather measurements, and then wait for all of the tasks to
: complete.

* schedgroup goroutines: naive prototype

.image naivecontext/goroutines.png 550 _

: The naive prototype uses time.AfterFunc, which spins up a goroutine to wait
: for the delay to elapse unless the timer is canceled. As a result, we end up
: with roughly 1 million parked goroutines waiting for a select statement to
: unblock.

* schedgroup goroutines: event-driven scheduler

Massive reduction in number of goroutines: ~1,000,000 to 3 (one is pprof!)

.image final/goroutines.png 450 _

: Since the final package only spins up a goroutine when necessary, we see
: a massive reduction in the number of running goroutines while waiting for the
: tasks to fire: from 1 million down to 3, and one of those is the pprof call!
: ---
: Of course, this task will still eventually execute one million goroutines once
: the delay for each task elapses, but it doesn't need to create them until exactly
: the time that they are needed.

* schedgroup heap in-use space: naive prototype

    total:          219.74MB

    runtime.mstart: 148.10MB
    time.After:      38.64MB
    other:           33.00MB

.image naivecontext/heapinuse.png 150 _

: Next, we will examine the in-use heap space. For this, I've chosen a flame graph
: because it provides a better view of the amount of memory a function call is
: using, in relation to other functions.
: ---
: runtime.mstart is the top heap space consumer here. This function creates new
: worker or machine threads, known as "Ms" in Go terminology. This means that the
: runtime had to create a large number of threads up-front in order to schedule all
: of the time.AfterFunc goroutines.
: ---
: The next consumer we see is space consumed by the state stored in time.AfterFunc.

* schedgroup heap in-use space: event-driven scheduler

Roughly 6x reduction in in-use heap bytes: 219.74MB to 36.50MB

    total:                     36.50MB

    schedgroup.(*Group).Delay: 30.50MB
    other:                      6.00MB

.image final/heapinuse.png 180 _

: The final package consumes much less memory while waiting for tasks to be scheduled
: because it doesn't need to create any threads for the goroutines to run until
: each task is ready to be scheduled. This results in roughly a 6x reduction in
: the in-use heap bytes while awaiting task completion. Do note that this
: number will inflate when the goroutines must actually be started, but not before.
: ---
: Group.Delay is the top memory consumer for this test, but it seems to consume
: a few less megabytes of memory than the time.AfterFunc equivalent.

* Summary

What did we learn?

- Proper measurement is essential to prove or disprove a performance hypothesis
- Focus on correctness and simplicity, then make iterative improvements
- Go runtime timers are very efficient, especially as of Go 1.14

.link https://golang.org/doc/go1.14#runtime

: In summary, I believe the ideal approach to tackling this sort of problem is to
: create a prototype, measure its performance, and use those measurements to
: challenge your assumptions and make incremental improvements.
: ---
: My primary goals are always correctness and readability. Performance optimization is a
: secondary goal that must be justified through benchmark and profile evidence.
: ---
: I was pleasantly surprised to see how well the Go runtime handles large numbers
: of timers, especially as of the improvements in Go 1.14. Maybe it wasn't strictly
: necessary to build out my own goroutine and timer management package, but I had
: a ton of fun doing it, and learned a lot along the way!

* More information

The final package is available on GitHub:

.link https://github.com/mdlayher/schedgroup

Special thanks to Egon Elbre for two initial prototypes which heavily influenced
the final design of the package.

.link https://twitter.com/egonelbre

Go issue: "runtime: make timers faster"

.link http://golang.org/issue/6239

CoreRAD: an extensible and observable IPv6 NDP router advertisement daemon

.link https://github.com/mdlayher/corerad

.link https://mdlayher.com/blog/corerad-a-new-ipv6-router-advertisement-daemon/

: You can check out the final package on GitHub at mdlayher/schedgroup.
: ---
: I'd like to again give special thanks to Egon Elbre for his help with two initial prototypes
: that inspired me to build out this package. I could not have done it without his help!
: ---
: You can learn more about the Go timer improvements on Go issue 6239, and if you'd like to check
: out CoreRAD, you can find the project on GitHub and an introduction on my blog.
