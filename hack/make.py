#!/usr/bin/env python

import sys
import subprocess
from os.path import expandvars

ROOT = expandvars('$GOPATH') + '/src/github.com/sadlil/mashable'

def call(cmd, stdin=None, cwd=ROOT):
    print(cmd)
    subprocess.call([expandvars(cmd)], shell=True, stdin=stdin, cwd=cwd)

def fmt():
    call('goimports -w .')
    call('go fmt ./...')


def vet():
    call('go vet ./...')

def deps():
    # specify git commit_hash to pin to specific version
    get_pkgs = [
        {'pkg': 'github.com/google/go-querystring/query'},
        {'pkg': 'github.com/stretchr/testify/assert'},
    ]
    for cfg in get_pkgs:
        call('go get -u ' + cfg['pkg'])
        if cfg.get('commit_hash', ''):
            call('git checkout ' + cfg['commit_hash'], cwd=expandvars('$GOPATH/src/' + cfg['pkg'].rstrip('/...')))
        if cfg.get('install', False):
            call('go install ./...', cwd=expandvars('$GOPATH/src/' + cfg['pkg'].rstrip('/...')))

def compile():
    call('go install ./...')

def test():
    call('go test')

def default():
    fmt()
    compile()
    test()

if __name__ == "__main__":
    if len(sys.argv) > 1:
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        default()