const data = require('./catalog.json');
const table = require('markdown-table');

function tableByType(type) {
    let t = [['Image', 'Args', 'Description', 'Source', 'Toolchain']];
    data.filter(r => r.type === type).forEach(r => {
        let desc = r.description;
        if (r.demo === true) {
            desc = '[Demo] ' + desc;
        }
        const source = `[Link](${r.source})`;
        let toolchain = ``;
        if (r.toolchain) {
          toolchain = `[SDK](${r.toolchain})`;
        }
        t.push([r.image, r.args, desc, source, toolchain]);
    });
    return table(t);
}

const README = `---
title: "Functions Catalog"
linkTitle: "Functions Catalog"
weight: 8
type: docs
description: >
    Catalog of Config Functions.
---

<!---
DO NOT EDIT. Generated by: "cd catalog; npm run gen-docs"
-->

This repository documents a catalog of functions implementing [Configuration Functions Specification][spec].

These functions can be implemented using any toolchain such as the [Typescript SDK][ts sdk] or [Golang SDK][go sdk].

## Sources

See [definition of source functions][source].

${tableByType('source')}

## Sinks

See [definition of sink functions][sink].

${tableByType('sink')}

## Validators

${tableByType('validator')}

## Generators

${tableByType('generator')}

## Transformers

${tableByType('transformer')}

## Miscellaneous

${tableByType('misc')}

[spec]: https://github.com/kubernetes-sigs/kustomize/blob/master/cmd/config/docs/api-conventions/functions-spec.md
[source]: ../../../../concepts/functions/#source-function
[sink]: ../../../../concepts/functions/#sink-function
[ts sdk]: ../../../producer/functions/ts
[go sdk]: ../../../producer/functions/golang`;

console.log(README);
