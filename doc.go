// Copyright 2018 Datawire. All rights reserved.

// Package k8scli provides command line interface functionality for
// github.com/ericchiang/k8s.
//
// There are 2 big reasons to prefer github.com/ericchiang/k8s over
// k8s.io/client-go:
//  1. It is much simpler and easy to understand
//  2. It has fewer dependencies, resulting in smaller binary sizes
//
// This package doesn't care about #2.
//
// Part of the complexity of client-go is that it mixes user-interface
// concerns with actual protocol concerns; ericchiang/k8s slims that
// down by only worrying about purely protocol concerns.  Which is
// great when you're writing something like a controller, but if
// you're writing a CLI program, then you lose a good chunk of
// functionality, like
//  - kubeconfig override flags
//  - correct kubeconfig file loading
//  - auth providers
//  - resource type shortcuts (eg "svc" instead of "service.v1.core")
//
// k8scli aims to implement these UI concerns on top of
// ericchiang/k8s, so that you don't need to write your code against
// the hard-to-use client-go.  For some of that functionality, k8scli
// uses client-go internally, so you don't get benefit #2 of using
// ericchiang/k8s.  But we believe that it's still worth-while in
// order to get benefit #1.
package k8scli
