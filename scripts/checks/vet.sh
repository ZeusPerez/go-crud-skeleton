#!/bin/bash
# vim: ai:ts=8:sw=8:noet
# Check for bugs
# Intended to be run from local machine or CI
set -eufo pipefail
IFS=$'\t\n'

# Check required commands are in place
command -v go >/dev/null 2>&1 || { echo 'please install go or use image that has it'; exit 1; }

go vet -composites=false -copylocks=false -printfuncs=Debug,Debugf,Debugln,Info,Infof,Infoln,Notice,Noticef,Noticeln,Error,Errorf,Errorln,Warning,Warningf,Warningln,Critical,Criticalf,Criticalln ./...

