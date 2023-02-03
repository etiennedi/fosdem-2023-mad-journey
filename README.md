# Our Mad Journey Of Building A Vector Database in Go

> This is the repo accompanying the 2023 FOSDEM talk of the same name.

## Recording

After the talk has happened, I'll put a link to the recording here

## Code

This repo contains the code that was used during the presentation. Mostly split
into two parts:

### Experiments

The `experiments` folder contains the code for the "before and after" experiments in the code used in the sections:

* Memory Allocations (28s -> 600ms)
* Delayed Encoding (14s -> 260ms -> 46ms)

To run the experiments, simply run `go test` in the `experiments` folder. To
isolate a specific experiment, and obtain cpu and mem profiles you can run
something like:

```
TODO
```

### Demo

Besides the experiment, you can also spin up a self-hosted demo instance,
similar to the one that was shown in the talk. 

To run the setup yourself, navigate to the `demo` folder. 

Export an OpenAI API key using 

```
TODO
```

Then spin up your setup using 

```
docker compose up -d
```

You can import the sample dataset using
```
python3 import.py
```

After the import has finished, you can query your instance using
```
TODO
```
