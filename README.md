Latenssi-Go
============

Latency measurement collector. Successor for [Latenssi](https://github.com/annttu/latenssi/).

Currently only icmp latency and loss collection using fping(6) is supported.

Probe
=====

Probe runs fping processes and send collected data points to collector over GRPC connection.

Collector
=========

Collects data points from Probe and stores those to time series database.
Currently InfluxDB is only supported time series database.

Installation
===========

  * Install go >= 1.17
  * Install fping
  * run make
  * Edit probe.yaml and collector.yaml

Usage
=====

bin/latenssi-probe
bin/latenssi-collector

License
=======

See LICENSE file.
