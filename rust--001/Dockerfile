FROM rust:latest

WORKDIR /app
COPY . /app 

RUN cargo build --release --all-targets

FROM debian:bookworm-slim as server1
COPY --from=builder /app/target/release/server1 server
ENTRYPOINT [ "/server" ]

FROM debian:bookworm-slim as server2
COPY --from=builder /app/target/release/server2 server
ENTRYPOINT [ "/server" ]

FROM debian:bookworm-slim as server3
COPY --from=builder /app/target/release/server3 server
ENTRYPOINT [ "/server" ]

FROM debian:bookworm-slim as server4
COPY --from=builder /app/target/release/server4 server
ENTRYPOINT [ "/server" ]