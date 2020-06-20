package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")
load("//:build/workspace.bzl", "RELEASE_VERSION")

#build package of toms go runtime.
pkg_tar(
    name = "gor-" + RELEASE_VERSION,
    extension = "tar.gz",
    srcs = [":gor"],
    mode = "0755",
)


go_binary(
    name = "gor",
    embed = [":main"],
)

go_library(
    name = "main",
    srcs = ["app.go"],
    importpath = "gor",
    deps = [
        ":server",
        ":config",
    ],
)

go_library(
    name = "server",
    srcs = [
        "pkg/server/server.go",
        "pkg/server/service.go",
    ],
    importpath = "gor/pkg/server",
    deps = [
        ":runtime_grpc",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes:go_default_library",
        "@org_golang_google_grpc//status:go_default_library",
    ],
)

go_library(
    name = "config",
    srcs = [
        "pkg/config/config.go",
    ],
    importpath = "gor/pkg/config",
    deps = [
        "@com_github_go_yaml//:go_default_library",
    ],
)

# gRPC
proto_library(
    name = "runtime_proto",
    srcs = ["build/protos/runtime.proto"],
    deps = ["@com_google_protobuf//:any_proto"],
)

go_proto_library(
    name = "runtime_grpc",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "gor/build/protos/runtime",
    proto = ":runtime_proto",
)
