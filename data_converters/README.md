

## Run

    # Run Temporal Server
    temporal server start-dev

    # Run Worker
    go run ./data_converters/...

    # Trigger Workflow
    tctl wf start --tq dummy-tq --wt Main
