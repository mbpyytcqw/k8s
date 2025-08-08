FROM golang:1.24 AS builder
ENV CGO_ENABLED=0
ARG BUILD_REF

COPY . /service

WORKDIR /service/api/services/sales
RUN go build -ldflags "-X main.build=${BUILD_REF}"


FROM alpine:3.22
RUN addgroup -g 1000 -S sales && \
    adduser -u 1000 -h /service -G sales -S sales
# COPY --from=build_sales --chown=sales:sales /service/api/tooling/admin/admin /service/admin
COPY --from=builder --chown=sales:sales /service/api/services/sales/sales /service/sales
WORKDIR /service
USER sales
CMD ["./sales"]
